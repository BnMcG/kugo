package main

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/bnmcg/kugo/authentication"
	"gopkg.in/yaml.v2"
)

// KubernetesClusterIdentityInformation Identify a server based on IP and CA information
type KubernetesClusterIdentityInformation struct {
	CertificateAuthorityData string `yaml:"certificate-authority-data"`
	Server                   string `yaml:"server"`
}

// KubernetesCluster represents configuration of an individual Kubernetes cluster
type KubernetesCluster struct {
	Name    string                               `yaml:"name"`
	Cluster KubernetesClusterIdentityInformation `yaml:"cluster"`
}

// KubernetesContextDetails hold cluster and user information for a context
type KubernetesContextDetails struct {
	Cluster string `yaml:"cluster"`
	User    string `yaml:"user"`
}

// KubernetesContext maps users to Kubernetes clusters
type KubernetesContext struct {
	Context KubernetesContextDetails `yaml:"context"`
	Name    string                   `yaml:"name"`
}

// KubernetesUser represents a user
type KubernetesUser struct {
	Name string                               `yaml:"name"`
	User authentication.KubernetesCredentials `yaml:"user"`
}

// KubernetesConfiguration represents .kube/config file
type KubernetesConfiguration struct {
	APIVersion     string              `yaml:"apiVersion"`
	CurrentContext string              `yaml:"current-context"`
	Kind           string              `yaml:"kind"`
	Clusters       []KubernetesCluster `yaml:"clusters"`
	Contexts       []KubernetesContext `yaml:"contexts"`
	Users          []KubernetesUser    `yaml:"users"`
	Preferences    interface{}         `yaml:"preferences"`
}

// LoadKubeconfig from file and return the parsed configuration
func LoadKubeconfig() (KubernetesConfiguration, error) {
	homeDir := os.Getenv("HOME")
	path := path.Join(homeDir, ".kube", "config")

	kubeconfigBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return KubernetesConfiguration{}, err
	}

	return ParseKubeconfig(kubeconfigBytes)
}

// ParseKubeconfig into struct
func ParseKubeconfig(config []byte) (KubernetesConfiguration, error) {
	kubeConfig := KubernetesConfiguration{}
	err := yaml.Unmarshal(config, &kubeConfig)

	if err != nil {
		return KubernetesConfiguration{}, err
	}

	return kubeConfig, nil
}

// WriteKubeconfig writes the given configuration to $HOME/.kube/config
func WriteKubeconfig(newConfig KubernetesConfiguration) error {
	serializedConfig, err := yaml.Marshal(newConfig)
	if err != nil {
		return err
	}

	home := os.Getenv("HOME")
	configPath := path.Join(home, ".kube", "config")
	err = ioutil.WriteFile(configPath, serializedConfig, 0644)
	if err != nil {
		return err
	}

	return nil
}
