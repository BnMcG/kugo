package authentication

import (
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/vault/api"
)

// VaultAuthenticator retrieves Kubernetes credentials from Hashicorp Vault
type VaultAuthenticator struct {
	Address            string
	PKIMount           string
	PKIRole            string
	KubernetesUsername string
	KubernetesTTL      string
}

// AuthenticateWithUsernamePassword retrieves credentials from Hashicorp Vault using the username/password authentication method
func (vaultAuthenticator *VaultAuthenticator) AuthenticateWithUsernamePassword(username string, password string) (KubernetesCredentials, error) {
	client, err := api.NewClient(&api.Config{
		Address: vaultAuthenticator.Address,
	})
	if err != nil {
		return KubernetesCredentials{}, err
	}

	token, err := exchangeUsernamePasswordForVaultToken(username, password, client)
	if err != nil {
		return KubernetesCredentials{}, err
	}

	client.SetToken(token)

	certificateRequestPayload := map[string]interface{}{
		"common_name": vaultAuthenticator.KubernetesUsername,
		"ttl":         vaultAuthenticator.KubernetesTTL,
	}

	certificateRequestPath := fmt.Sprintf("%s/issue/%s", vaultAuthenticator.PKIMount, vaultAuthenticator.PKIRole)

	certificateSecret, err := client.Logical().Write(certificateRequestPath, certificateRequestPayload)
	if err != nil {
		return KubernetesCredentials{}, err
	}

	PEMCertificateAsString := certificateSecret.Data["certificate"].(string)
	RSAPrivateKeyAsString := certificateSecret.Data["private_key"].(string)

	return KubernetesCredentials{
		ClientCertificateData: base64.StdEncoding.EncodeToString([]byte(PEMCertificateAsString)),
		ClientKeyData:         base64.StdEncoding.EncodeToString([]byte(RSAPrivateKeyAsString)),
	}, nil
}

func exchangeUsernamePasswordForVaultToken(username string, password string, api *api.Client) (string, error) {
	loginPath := fmt.Sprintf("auth/userpass/login/%s", username)
	payload := map[string]interface{}{
		"password": password,
	}

	secret, err := api.Logical().Write(loginPath, payload)
	if err != nil {
		return "", err
	}

	return secret.Auth.ClientToken, nil
}
