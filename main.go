package main

import "github.com/skybet/cali"

func main() {
	cli := cali.Cli("lucli")
	cli.SetShort("Example CLI tool")
	cli.SetLong("A nice long description of what your tool actually does")

	terraform := cli.Command("terraform [command]")
	terraform.SetShort("Run Terraform in an ephemeral container")
	terraform.SetLong(`Starts a container for Terraform and attempts to run it against your code. There are two choices for code source; a local mount, or directly from a git repo.

Examples:

  To build the contents of the current working directory using my_account as the AWS profile from the shared credentials file on this host.
  # lucli terraform plan -p my_account

  Any addtional flags sent to the terraform command come after the --, e.g.
  # lucli terraform plan -- -state=environments/test/terraform.tfstate -var-file=environments/test/terraform.tfvars
  # lucli terraform -- plan -out plan.out
`)
	terraform.Flags().StringP("profile", "p", "default", "Profile to use from the AWS shared credentials file")
	terraform.BindFlags()

	terraformTask := terraform.Task("hashicorp/terraform:0.11.1")
	terraformTask.SetInitFunc(func(t *cali.Task, args []string) {
		t.AddEnv("AWS_PROFILE", cli.FlagValues().GetString("profile"))
	})

	cli.Start()
}
