package main

import (
	"context"
	"log"
	"os"

	workflowDefinitions "github.com/roshbhatia/scaffold/src/cmd/std/workflow"
	"github.com/roshbhatia/scaffold/src/shared/config"
	"go.temporal.io/sdk/client"
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
	
	we, err := c.ExecuteWorkflow(context.Background(), options, workflowDefinitions.KubernetesDeploymentWorkflow, configPath)
	if err != nil {
		log.Fatalln("Failed to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
