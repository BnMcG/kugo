package configuration

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// KugoConfiguration is the wrapper configuration
type KugoConfiguration struct {
	VaultAddress  string `yaml:"vault_address"`
	VaultUsername string `yaml:"vault_username"`
	VaultPassword string `yaml:"vault_password"`
	VaultPKIRole  string `yaml:"vault_pki_role"`
	VaultPKIMount string `yaml:"vault_pki_mount"`

	KubernetesPKITTL string `yaml:"kubernetes_pki_ttl"`
}

// LoadConfiguration from $HOME/.kugo.yaml
func LoadConfiguration() (KugoConfiguration, error) {
	homeDirectory := os.Getenv("HOME")
	configurationPath := path.Join(homeDirectory, ".kugo.yaml")
	configurationBytes, err := ioutil.ReadFile(configurationPath)
	if err != nil {
		return KugoConfiguration{}, err
	}

	configuration := KugoConfiguration{}
	err = yaml.Unmarshal(configurationBytes, &configuration)
	if err != nil {
		return KugoConfiguration{}, err
	}

	return configuration, nil
}
