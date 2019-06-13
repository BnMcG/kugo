package main

import "testing"

var exampleSingleClusterConfiguration = `
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: testCertificateAuthority
    server: https://127.0.0.1:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate-data: testClientCertificateData
    client-key-data: testClientKeyData`

var exampleMultipleClusterConfiguration = `
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: testCertificateAuthority
    server: https://127.0.0.1:6443
  name: kubernetes
- cluster:
    certificate-authority-data: testCertificateAuthority2
    server: https://127.0.0.2:6443
  name: kubernetes_two
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
- context:
    cluster: kubernetes2
    user: kubernetes-admin2
  name: kubernetes-admin2@kubernetes2
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate-data: testClientCertificateData
    client-key-data: testClientKeyData
- name: kubernetes-admin2
  user:
    client-certificate-data: testClientCertificateData2
    client-key-data: testClientKeyData2`

func TestAPIVersionParsing(t *testing.T) {
	config, err := ParseKubeconfig([]byte(exampleSingleClusterConfiguration))
	if err != nil {
		t.Error(err)
	}

	if config.APIVersion != "v1" {
		t.Errorf("Incorrect API version parsed!")
	}
}

func TestKindParsing(t *testing.T) {
	config, err := ParseKubeconfig([]byte(exampleSingleClusterConfiguration))
	if err != nil {
		t.Error(err)
	}

	if config.Kind != "Config" {
		t.Errorf("Incorrect kind parsed!")
	}
}

func TestCurrentContextParsing(t *testing.T) {
	config, err := ParseKubeconfig([]byte(exampleSingleClusterConfiguration))
	if err != nil {
		t.Error(err)
	}

	if config.CurrentContext != "kubernetes-admin@kubernetes" {
		t.Errorf("Incorrect current context parsed!")
	}
}

func TestSingleClusterParsing(t *testing.T) {
	config, err := ParseKubeconfig([]byte(exampleSingleClusterConfiguration))
	if err != nil {
		t.Error(err)
	}

	if len(config.Clusters) != 1 {
		t.Error("Incorrect number of clusters parsed!")
	}

	validateCluster(config.Clusters[0], t, "testCertificateAuthority", "https://127.0.0.1:6443", "kubernetes")
}

func TestMultipleClusterParsing(t *testing.T) {
	config, err := ParseKubeconfig([]byte(exampleMultipleClusterConfiguration))
	if err != nil {
		t.Error(err)
	}

	if len(config.Clusters) != 2 {
		t.Error("Incorrect number of clusters parsed!")
	}

	validateCluster(config.Clusters[0], t, "testCertificateAuthority", "https://127.0.0.1:6443", "kubernetes")
	validateCluster(config.Clusters[1], t, "testCertificateAuthority2", "https://127.0.0.2:6443", "kubernetes_two")
}

func TestSingleContextParsing(t *testing.T) {
	config, err := ParseKubeconfig([]byte(exampleSingleClusterConfiguration))
	if err != nil {
		t.Error(err)
	}

	if len(config.Contexts) != 1 {
		t.Error("Incorrect number of clusters parsed!")
	}

	validateContext(config.Contexts[0], t, "kubernetes", "kubernetes-admin", "kubernetes-admin@kubernetes")
}

func TestMultipleContextParsing(t *testing.T) {
	config, err := ParseKubeconfig([]byte(exampleMultipleClusterConfiguration))
	if err != nil {
		t.Error(err)
	}

	if len(config.Contexts) != 2 {
		t.Error("Incorrect number of contexts parsed!")
	}

	validateContext(config.Contexts[0], t, "kubernetes", "kubernetes-admin", "kubernetes-admin@kubernetes")
	validateContext(config.Contexts[1], t, "kubernetes2", "kubernetes-admin2", "kubernetes-admin2@kubernetes2")
}

func TestSingleUserParsing(t *testing.T) {
	config, err := ParseKubeconfig([]byte(exampleSingleClusterConfiguration))
	if err != nil {
		t.Error(err)
	}

	if len(config.Users) != 1 {
		t.Error("Incorrect number of users parsed!")
	}

	validateUser(config.Users[0], t, "kubernetes-admin", "testClientCertificateData", "testClientKeyData")
}

func TestMultipleUserParsing(t *testing.T) {
	config, err := ParseKubeconfig([]byte(exampleMultipleClusterConfiguration))
	if err != nil {
		t.Error(err)
	}

	if len(config.Users) != 2 {
		t.Error("Incorrect number of users parsed!")
	}

	validateUser(config.Users[0], t, "kubernetes-admin", "testClientCertificateData", "testClientKeyData")
	validateUser(config.Users[1], t, "kubernetes-admin2", "testClientCertificateData2", "testClientKeyData2")
}

func validateCluster(cluster KubernetesCluster, t *testing.T, expectedCAData string, expectedServer string, expectedName string) {
	if cluster.Cluster.CertificateAuthorityData != expectedCAData {
		t.Error("Incorrect cluster certificate authority data parsed!")
	}

	if cluster.Cluster.Server != expectedServer {
		t.Error("Incorrect cluster server data parsed!")
	}

	if cluster.Name != expectedName {
		t.Error("Incorrect cluster name parsed!")
	}
}

func validateContext(context KubernetesContext, t *testing.T, expectedCluster string, expectedUser string, expectedName string) {
	if context.Context.Cluster != expectedCluster {
		t.Error("Incorrect cluster name context data parsed!")
	}

	if context.Context.User != expectedUser {
		t.Error("Incorrect user context data parsed!")
	}

	if context.Name != expectedName {
		t.Error("Incorrect name context data parsed!")
	}
}

func validateUser(user KubernetesUser, t *testing.T, expectedName string, expectedClientCertificateData string, expectedClientKeyData string) {
	if user.Name != expectedName {
		t.Error("Incorrect user name data parsed!")
	}

	if user.User.ClientCertificateData != expectedClientCertificateData {
		t.Error("Incorrect user certificate data parsed!")
	}

	if user.User.ClientKeyData != expectedClientKeyData {
		t.Error("Incorrect user key data parsed!")
	}
}
