# A simple buffalo API application

A simple Golang application which
  -  exposes API end points
  -  support gRPC
  -  built with buffalo framework
  -  basic authentication
  -  postgres database


## Database Setup

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

You will also need to make sure that **you** start/install the database of your choice. Buffalo **won't** install and start postgres for you.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started postgres, now Buffalo can create the databases in that file for you:

	$ buffalo db create -a

### Run migrations

    $ buffalo pop migrate

### Run Seed migrations

    $ buffalo task db:seed:wc

### Get the routes

    $ buffalo routes

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev


## Sample http request

    $  curl localhost:3000 -H 'ClientID: user001' -H 'PassKey: pass001'

### gRPC Server

   This application is running a simple gRPC server in port 50051.

### Run gRPC client to test the gRPC server

    $ cd grpc-client
    $ go run run-client.go

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.