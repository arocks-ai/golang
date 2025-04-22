package main

import (
	"log"

	//"github.com/temporalio/samples-go/helloworld"
	"temporal-hello/app"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	log.Println("Starting worker...")

	// Create a Temporal client
	temporalClient, err := client.Dial(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "default",
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer temporalClient.Close()

	worker1 := worker.New(temporalClient, "hello-world-task-queue", worker.Options{})

	worker1.RegisterWorkflow(app.HelloWorkflow) // HelloWorkflow Workflow function in app package
	worker1.RegisterActivity(app.HelloActivity) // HelloActivity Activity function in app package

	err = worker1.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

	log.Println("Ending worker...")

}
