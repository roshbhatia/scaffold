package workflows

import (
	"github.com/roshbhatia/scaffold/pkg/config"
	kubernetesShim "github.com/roshbhatia/scaffold/pkg/kubernetes"
)

type WorkflowManager struct {
	ConfigReader     config.ConfigReader
	KubernetesClient kubernetesShim.KubernetesClient
}
