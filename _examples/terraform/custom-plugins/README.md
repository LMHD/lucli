# Custom Terraform Plugins

Here's an example of using a custom Terraform plugin

Plugin URLs are specified as an array in your local lucli config file, i.e.

```
terraform:
  plugins:
   urls:
   - https://example.com/foo/bar
```

Plugins are downloaded to the `terraform.d/plugins/linux_amd64` directory as
part of the init function, when running `lucli terraform init`, prior to
starting the Docker container.

As it happens, the Terraform code in this example directory doesn't work.

It uses the https://github.com/Mastercard/terraform-provider-restapi, which
makes assumptions about the format of the API you're calling, that do not hold
true for https://whoami.lmhd.me/name

But you can at least see that the custom plugin works fine.
