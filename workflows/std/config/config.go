package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	KubernetesDeploymentName string `yaml:"kubernetes_deployment_name"`
	KubernetesDefaultReplicas int32 `yaml:"kubernetes_default_replicas"`
	KubernetesNamespace string `yaml:"kubernetes_namespace"`
	DockerImageURI string `yaml:"docker_image_uri"`
}

func ReadConfig(filename string) (Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
