### Steps to run this program:

1) Project setup
Reference - Run a [Temporal service](https://github.com/temporalio/samples-go/tree/main/#how-to-use).

```
$ mkdir goproject
$ cd goproject
$ go mod init goproject/app 	# Initialize a Go project in that directory:
$ go get go.temporal.io/sdk	# Install the Temporal SDK with go get:
```

1b)  Start the server 
```
temporal server start-dev
```

2) Start the worker
```
go run helloworld/worker/main.go
```

3) Start the client application
```
go run helloworld/starter/main.go
```
