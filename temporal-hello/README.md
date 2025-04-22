### Steps to run this program:

1) Project setup

```
$ mkdir temporal-hello
$ cd temporal-hello
$ go mod init temporal-hello/app 	# Initialize a Go project 
$ go get go.temporal.io/sdk	        # Install/download the Temporal SDK 
```
Create the required files in the above directory.

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

3b) View the Workflow status & History from the Web UI console at http://localhost:8233/

4) Simulate the error

Return err message for activity return code
```
 // return "Hello " + name + "!", nil
 return "Hello " + name + "!", fmt.Errorf("Error occurred in HelloActivity") // Simulate an error
```

4b) Web UI console will show activity still running and events shows the failures