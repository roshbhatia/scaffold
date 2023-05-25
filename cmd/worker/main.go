package main

import (
	"log"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"github.com/roshbhatia/scaffold/pkg/config"
	kubernetesShim "github.com/roshbhatia/scaffold/pkg/kubernetes"
	workflowDefinitions "github.com/roshbhatia/scaffold/pkg/workflows"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Failed to create temporal client", err)
	}
	defer c.Close()

	// This worker hosts both Worker and Activity functions
	w := worker.New(c, config.TaskQueueName, worker.Options{})

	// Create Kubernetes client
	kubeconfig := os.Getenv("KUBECONFIG")
	kubeConfigObj, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalln("Failed to build config", err)
	}
	clientSet, err := kubernetes.NewForConfig(kubeConfigObj)
	if err != nil {
		log.Fatalln("Failed to create client", err)
	}

	// Create dependencies
	configReader := config.DefaultConfigReader{}
	kubeClient := kubernetesShim.KubernetesClient{
		Client: clientSet,
	}

	// Register workflow with dependencies
	w.RegisterWorkflow(func(ctx workflow.Context, configFilePath string) error {
		return workflowDefinitions.KubernetesDeploymentWorkflow(ctx, &configReader, kubeClient, configFilePath)
	})

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Failed to start worker", err)
	}
}
