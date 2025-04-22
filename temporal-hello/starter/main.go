package main

import (
	"context"
	"log"
	"temporal-hello/app"

	"go.temporal.io/sdk/client"
)

func main() {
	log.Println("Starting starter::main ...")

	// Create a Temporal client
	temporalClient, err := client.Dial(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "default",
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer temporalClient.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "hello-world-workflow",
		TaskQueue: "hello-world-task-queue",
	}

	workExecution, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, app.HelloWorkflow, "World")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Workflow started with ID:", workExecution.GetID(), " RunID:", workExecution.GetRunID())

	// Synchronously Wait for the workflow to complete
	var result string
	err = workExecution.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

	log.Println("Workflow result:", result)
	log.Println("Ending starter::main ...")
}
