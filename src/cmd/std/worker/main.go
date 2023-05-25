package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	workflowDefinitions "github.com/roshbhatia/scaffold/src/cmd/std/workflow"
	"github.com/roshbhatia/scaffold/src/shared/config"
	kubernetesShim "github.com/roshbhatia/scaffold/src/shared/kubernetes"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Failed to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, config.TaskQueueName, worker.Options{})
	w.RegisterWorkflow(workflowDefinitions.KubernetesDeploymentWorkflow)
	w.RegisterActivity(kubernetesShim.CreateDeployment)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Failed to start worker", err)
	}
}
