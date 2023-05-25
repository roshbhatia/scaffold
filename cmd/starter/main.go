package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/roshbhatia/scaffold/pkg/config"
	"github.com/roshbhatia/scaffold/pkg/kubernetes"
	"github.com/roshbhatia/scaffold/pkg/workflows"
	"go.temporal.io/sdk/client"
	"k8s.io/klog/v2"
)

func main() {
	// Initialize klog flags
	klog.InitFlags(nil)
	flag.Parse()

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

	// Read the configuration file path from the environment variable
	configPath := os.Getenv("SCAFFOLD_CONFIG_PATH")
	if configPath == "" {
		log.Fatalln("Failed to get SCAFFOLD_CONFIG_PATH from environment variables")
	}

	// Create dependencies
	configReader := config.ConfigReader{}
	kubeClient := kubernetes.KubernetesClient{}

	// Create the workflow manager with dependencies
	workflowManager := &workflows.WorkflowManager{
		ConfigReader:     configReader,
		KubernetesClient: kubeClient,
	}

	// Start a workflow execution
	we, err := c.ExecuteWorkflow(context.Background(), options, workflowManager.KubernetesDeploymentWorkflow, configPath)
	if err != nil {
		log.Fatalln("Failed to execute workflow", err)
	}

	klog.Info("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
