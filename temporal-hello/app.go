package app

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
)

func HelloWorkflow(ctx workflow.Context, name string) (string, error) {

	// Set up activity options
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	// Get the logger from the workflow context
	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorkflow started ...", "name", name)

	// Execute the activity
	var result string
	err := workflow.ExecuteActivity(ctx, HelloActivity, name).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	logger.Info("HelloWorkflow Ended ...", "name", name)
	return result, nil
}

// Activity function
// function take context and string as input
//           and returns string, error

func HelloActivity(ctx context.Context, name string) (string, error) {
	// Simulate some work
	return "Hello " + name + "!", nil
	//return "Hello " + name + "!", fmt.Errorf("Error occurred in HelloActivity") // Simulate an error
}
