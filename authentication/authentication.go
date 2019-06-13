package authentication

// Authenticator handles authenticating with an external identity provider and retrieving credentials for Kubernetes
type Authenticator interface {
	AuthenticateWithUsernamePassword(username string, password string) (KubernetesCredentials, error)
}

// KubernetesCredentials represents credentials a user uses to authenticate to a Kubernetes cluster
type KubernetesCredentials struct {
	ClientCertificateData string `yaml:"client-certificate-data"`
	ClientKeyData         string `yaml:"client-key-data"`
}
