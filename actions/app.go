package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/mw-forcessl"
	"github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	"github.com/gobuffalo/mw-contenttype"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
	"github.com/simple_buffalo/models"

	"github.com/sirupsen/logrus"
	"net"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"fmt"
	"golang.org/x/net/context"
	pb "github.com/simple_buffalo/proto"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")

const (
	HEADER_KEY_CLIENT_ID = "ClientID"
	HEADER_KEY_PASS_KEY  = "PassKey"
)

var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_simple_buffalo_session",
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Set the request content type to JSON
		app.Use(contenttype.Set("application/json"))

		app.Use(AuthMiddleware)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		app.GET("/", HomeHandler)

		go func() {
			fmt.Println("Starting teh gRPC server")
			startGRPCServer()
		}()
	}

	return app
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Recieved a request in gRPC")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func AuthMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		clientID := c.Request().Header.Get(HEADER_KEY_CLIENT_ID)
		passKey := c.Request().Header.Get(HEADER_KEY_PASS_KEY)

		err := NewClientAuthentication().Authenticate(clientID, passKey)
		if err != nil {
			logrus.Error("Authentication failed!!!")
			return c.Error(401, err)
		}

		err = next(c)
		// do some work after calling the next handler
		return err
	}
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
