# Kugo
Kugo is a wrapper for kubectl supporting additional authentication options. It was written with the purpose of allowing authentication
to a Kubernetes Cluster using Hashicorp Vault.

## Configuring
Kugo is configured using a YAML file. This file must be placed at $HOME/.kugo.yaml. An example is below:

### Hashicorp Vault authentication with username/password
```yaml
vault_address: https://vault:8443
vault_username: kugo
vault_password: password
vault_pki_role: kugo-pki
vault_pki_mount: pki
kubernetes_pki_ttl: 1d 
```