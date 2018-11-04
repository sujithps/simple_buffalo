package main

import (
	"log"

	"os"

	"github.com/gobuffalo/envy"
	"github.com/simple_buffalo/actions"
	"github.com/sirupsen/logrus"
)

// main is the starting point for your Buffalo application.
// You can feel free and add to this `main` method, change
// what it does, etc...
// All we ask is that, at some point, you make sure to
// call `app.Serve()`, unless you don't want to start your
// application that is. :)

func main() {
	configureLog()

	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}

func configureLog() {
	var LOG_LEVEL = envy.Get("LOG_LEVEL", "info")

	level, err := logrus.ParseLevel(LOG_LEVEL)
	if err != nil {
		panic(err.Error())
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(level)
}

/*
# Notes about `main.go`

## SSL Support

We recommend placing your application behind a proxy, such as
Apache or Nginx and letting them do the SSL heavy lifting
for you. https://gobuffalo.io/en/docs/proxy

## Buffalo Build

When `buffalo build` is run to compile your binary, this `main`
function will be at the heart of that binary. It is expected
that your `main` function will start your application using
the `app.Serve()` method.

*/
