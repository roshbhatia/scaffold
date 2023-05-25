package main

import (
	"log"

	"github.com/roshbhatia/scaffold/workflows/std/config"
	"github.com/roshbhatia/scaffold/workflows/std/shared"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Failed to create client", err)
	}
	defer c.Close()

	// Create a worker that listens on the TaskQueue and hosts the workflow and activity implementations.
	w := worker.New(c, config.TaskQueueName, worker.Options{})

	// Register the workflow and activity with the worker.
	w.RegisterWorkflow(shared.KubernetesDeploymentWorkflow)
	w.RegisterActivity(shared.DeployKubernetes)

	// Start listening to the TaskQueue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Failed to start worker", err)
	}
}
