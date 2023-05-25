package workflow

import (
	"time"

	"github.com/roshbhatia/scaffold/src/shared/config"
	kubernetesShim "github.com/roshbhatia/scaffold/src/shared/kubernetes"
	"go.temporal.io/sdk/workflow"
)


func KubernetesDeploymentWorkflow(ctx workflow.Context, configFilePath string) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    time.Minute * 10, // Timeout for the activity
		HeartbeatTimeout:       time.Second * 2,  // Heartbeat interval
		ScheduleToCloseTimeout: time.Minute * 10, // Total timeout that includes retries
	})

	// Read the config file
	cfg, err := config.ReadConfig(configFilePath)
	if err != nil {
		workflow.GetLogger(ctx).Error("Failed to read configuration", "Error", err)
		return err
	}

	// Deploy to Kubernetes using the config
	err = workflow.ExecuteActivity(ctx, kubernetesShim.CreateDeployment, cfg.KubernetesDeploymentName, cfg.DockerImageURI, cfg.KubernetesDefaultReplicas).Get(ctx, nil)
	if err != nil {
		workflow.GetLogger(ctx).Error("Failed to deploy Kubernetes", "Error", err)
		return err
	}

	return nil
}
