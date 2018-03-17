package cmd

import (
	"fmt"

	"github.com/skybet/cali"
)

func init() {

	// Default image name/version
	imageName := "hashicorp/terraform"
	imageVersion := "latest"

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
	command.Flags().StringP("profile", "p", "default", "Profile to use from the AWS shared credentials file")
	command.Flags().StringP("terraform-version", "v", imageVersion, "Version of image to use")
	command.BindFlags()

	// Set default image for Run function
	task := command.Task(fmt.Sprintf("%s:%s", imageName, imageVersion))

	// TODO: some amazing witchcraft, with sidekick containers, providing stuff like awscli

	// Init function, set profile, and image version
	task.SetInitFunc(func(t *cali.Task, args []string) {
		t.AddEnv("AWS_PROFILE", cli.FlagValues().GetString("profile"))
		t.SetImage(fmt.Sprintf("%s:%s", imageName, cli.FlagValues().GetString("terraform-version")))
	})

}
