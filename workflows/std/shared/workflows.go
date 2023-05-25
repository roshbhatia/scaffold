package shared

import (
	"context"
	"os"
	"time"

	"github.com/roshbhatia/scaffold/workflows/std/config"
	"go.temporal.io/sdk/workflow"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func DeployKubernetes(ctx context.Context, deploymentName string, imageName string, replicas int32) error {
	kubeconfig := os.Getenv("KUBECONFIG")
	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig}
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})
	restConfig, err := config.ClientConfig()
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": deploymentName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deploymentName,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  deploymentName,
							Image: imageName,
						},
					},
				},
			},
		},
	}

	_, err = clientset.AppsV1().Deployments("default").Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

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
	err = workflow.ExecuteActivity(ctx, DeployKubernetes, cfg.KubernetesDeploymentName, cfg.DockerImageURI, cfg.KubernetesDefaultReplicas).Get(ctx, nil)
	if err != nil {
		workflow.GetLogger(ctx).Error("Failed to deploy Kubernetes", "Error", err)
		return err
	}

	return nil
}
