# Kugo
Kugo is a wrapper for kubectl supporting additional authentication options. It was written with the purpose of allowing authentication
to a Kubernetes Cluster using Hashicorp Vault.

## Configuring
Kugo is configured using a YAML file. This file must be placed at `$HOME/.kugo.yaml`. An example is below:

### Hashicorp Vault authentication with username/password
```yaml
vault_address: https://vault:8443
vault_username: kugo
vault_password: password
vault_pki_role: kugo-pki
vault_pki_mount: pki
kubernetes_pki_ttl: 1d 
```

## Wrapping other executables
Kugo may also wrap around other executables in the Kubernetes ecosystem. Some examples would be Helm and Telepresence. By wrapping around other applications, Kugo can also refresh your Kubernetes credentials before
executing these tools. In order to wrap around other applications, just pass the `-exectuable` flag, like so:

```
kugo -executable=helm install stable/nginx
kugo -executable=telepresence --namespace test
```

Note: The executable flag must be passed before the arguments you wish to pass through to the wrapped application!

## Shell aliases
### Fish
You can setup an alias in your Fish shell in order to execute Kugo instead of the wrapped application. Your alias may either overwrite the existing name, or use a new name. Examples are below:

```
alias k "/home/user/kugo"
alias telepresence "/home/user/kugo -executable=telepresence"
alias helm "/home/user/kugo -executable=helm"
```