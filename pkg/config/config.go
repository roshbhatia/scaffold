package config

type Config struct {
	KubernetesDeploymentName string `yaml:"kubernetes_deployment_name"`
	KubernetesDefaultReplicas int32 `yaml:"kubernetes_default_replicas"`
	KubernetesNamespace string `yaml:"kubernetes_namespace"`
	DockerImageURI string `yaml:"docker_image_uri"`
}
