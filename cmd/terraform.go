package cmd

import (
	"fmt"

	"github.com/lmhd/lucli/creds"
	"github.com/skybet/cali"
)

func init() {

	command := cli.NewCommand("terraform [command]")

	command.SetShort("Run Terraform in an ephemeral container")

	command.SetLong(`Starts a container for Terraform and attempts to run it against your code. There are two choices for code source; a local mount, or directly from a git repo.

Examples:

  To build the contents of the current working directory using my_account as the AWS profile from the shared credentials file on this host.
  # lucli terraform plan -p my_account

  Any addtional flags sent to the terraform command come after the --, e.g.
  # lucli terraform plan -- -state=environments/test/terraform.tfstate -var-file=environments/test/terraform.tfvars
  # lucli terraform -- plan -out plan.out
`)

	// Set default image for Run function
	imageName := "hashicorp/terraform"
	imageVersion := "latest"

	command.Flags().StringP("terraform-version", "v", imageVersion, "Version of image to use")
	command.BindFlags()

	task := command.Task(fmt.Sprintf("%s:%s", imageName, imageVersion))

	// Init function, set profile, and image version
	task.SetInitFunc(func(t *cali.Task, args []string) {
		t.SetImage(fmt.Sprintf("%s:%s", imageName, cli.FlagValues().GetString("terraform-version")))

		_ = creds.BindAWS(t, args)
	})

}
