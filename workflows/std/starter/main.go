package main

import (
	"context"
	"log"
	"os"

	"go.temporal.io/sdk/client"

	"github.com/roshbhatia/scaffold/workflows/std/config"
	"github.com/roshbhatia/scaffold/workflows/std/shared"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Failed to create client", err)
	}
	defer c.Close()

	// Define the workflow execution parameters
	options := client.StartWorkflowOptions{
		ID:        "scaffold_workflow",
		TaskQueue: config.TaskQueueName,
	}

	// Start a workflow execution.
	// The configuration file path is read from the environment variable
	configPath := os.Getenv("SCAFFOLD_CONFIG_PATH")
	if configPath == "" {
		log.Fatalln("Failed to get SCAFFOLD_CONFIG_PATH from environment variables")
	}
	
	we, err := c.ExecuteWorkflow(context.Background(), options, shared.KubernetesDeploymentWorkflow, configPath)
	if err != nil {
		log.Fatalln("Failed to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
