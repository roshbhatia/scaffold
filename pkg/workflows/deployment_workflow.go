package workflows

import (
	"time"

	"go.temporal.io/sdk/workflow"
)


func (w *WorkflowManager) KubernetesDeploymentWorkflow(ctx workflow.Context, configFilePath string) error {
	cfg, err := w.ConfigReader.ReadConfig(configFilePath)
	// Read the config file
	if err != nil {
		workflow.GetLogger(ctx).Error("Failed to read configuration", "Error", err)
		return err
	}

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    time.Minute * 10, // Timeout for the activity
		HeartbeatTimeout:       time.Second * 2,  // Heartbeat interval
		ScheduleToCloseTimeout: time.Minute * 10, // Total timeout that includes retries
	})

	// Deploy to Kubernetes using the config
	err = workflow.ExecuteActivity(ctx, w.KubernetesClient.CreateDeployment, cfg.KubernetesDeploymentName, cfg.DockerImageURI, cfg.KubernetesDefaultReplicas, cfg.KubernetesNamespace).Get(ctx, nil)
	if err != nil {
		workflow.GetLogger(ctx).Error("Failed to deploy Kubernetes", "Error", err)
		return err
	}

	return nil
}
