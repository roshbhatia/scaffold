package main

import (
	"context"
	"log"

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
		ID:        "kubernetes_deployment_workflow",
		TaskQueue: config.TaskQueueName,
	}

	// Start a workflow execution. 
	we, err := c.ExecuteWorkflow(context.Background(), options, shared.KubernetesDeploymentWorkflow, config.DeploymentName, config.ImageName, config.Replicas)
	if err != nil {
		log.Fatalln("Failed to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
