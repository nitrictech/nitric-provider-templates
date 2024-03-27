Building the extension provider can be done with the following command:

```bash
make install
```

This will build the runtime provider and the deployment provider, packaging them together and saving it to `$HOME/.nitric/providers/custom/extension-0.0.1`.

To use the custom extension you can use the following stack configuration file. It requires you fill in digital ocean tokens to deploy your spaces bucket.

- [digital_ocean_token](https://cloud.digitalocean.com/account/api/tokens)
- [spaces_key](https://cloud.digitalocean.com/account/api/spaces)
- [spaces_secret](https://cloud.digitalocean.com/account/api/spaces)

```yaml
provider: custom/extension@0.0.1
region: us-east-1
token: `digital_ocean_token`
spaces:
  region: nyc1
  key: `spaces_key`
  secret: `spaces_secret`
```
