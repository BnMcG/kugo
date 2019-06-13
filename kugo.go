package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/bnmcg/kugo/authentication"
	"github.com/bnmcg/kugo/configuration"
)

func main() {
	configuration, err := configuration.LoadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	// Parse existing k8s configuration
	kubeconfig, err := LoadKubeconfig()
	if err != nil {
		log.Fatal(err)
	}

	// Get current context
	var currentContext KubernetesContext
	for _, context := range kubeconfig.Contexts {
		if context.Name == kubeconfig.CurrentContext {
			currentContext = context
			break
		}
	}

	// Get current user
	var currentUser KubernetesUser
	var currentUserIndex int
	for index, user := range kubeconfig.Users {
		if user.Name == currentContext.Context.User {
			currentUser = user
			currentUserIndex = index
			break
		}
	}

	currentCertificate, err := DecodeBase64EncodedPEMCertificate(currentUser.User.ClientCertificateData)
	if err != nil {
		log.Fatal(err)
	}

	if CertificateHasExpired(currentCertificate) {
		authenticator := authentication.VaultAuthenticator{
			Address:            configuration.VaultAddress,
			PKIMount:           configuration.VaultPKIMount,
			PKIRole:            configuration.VaultPKIRole,
			KubernetesUsername: currentUser.Name,
			KubernetesTTL:      configuration.KubernetesPKITTL,
		}
		newCredentials, err := authenticator.AuthenticateWithUsernamePassword(configuration.VaultUsername, configuration.VaultPassword)
		if err != nil {
			log.Fatal(err)
		}

		kubeconfig.Users[currentUserIndex].User = newCredentials
		err = WriteKubeconfig(kubeconfig)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("[Kugo] Refreshed Kubernetes credentials...\n")
	}

	cmd := exec.Command("kubectl", os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
