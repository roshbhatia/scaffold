package kubernetes

import (
	"context"

	"k8s.io/client-go/kubernetes"
)

type KubernetesClientInterface interface {
	CreateDeployment(ctx context.Context, deploymentName string, imageName string, replicas int32, namespace string) error
}

type KubernetesClient struct {
	Client *kubernetes.Clientset
}
