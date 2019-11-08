package cmd

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"github.com/lmhd/lucli/creds"
	"github.com/lmhd/lucli/lib"
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
		finalImageName := fmt.Sprintf("%s:%s", imageName, cli.FlagValues().GetString("terraform-version"))
		t.SetImage(finalImageName)

		task.AddEnv("AWS_PROFILE", cli.FlagValues().GetString("aws-profile"))
		_ = creds.BindAWS(t, args)

		// For terraform init only, download custom plugins, if any
		if len(args) > 0 && args[0] == "init" {
			if cli.FlagValues().IsSet("terraform.plugins.urls") {
				pluginURLs := cli.FlagValues().GetStringSlice("terraform.plugins.urls")
				log.Infof("Using %v custom plugins: %v", len(pluginURLs), pluginURLs)

				err := downloadTerraformPlugins(pluginURLs)
				if err != nil {
					log.Fatalf("Error downloading plugin: %s", err)
				}

				// Apply workaround to Terraform image
				err = fixTerraformImageForCustomPlugins(task, finalImageName)
				if err != nil {
					log.Fatalf("Error hacking the Terraform image: %s", err)
				}
			}
		}
	})

}

// downloadTerraformPlugins downloads plugins from a slice of URLs
func downloadTerraformPlugins(pluginURLs []string) error {
	pluginDir := "terraform.d/plugins/linux_amd64"

	// TODO: In future, could add some custom handling for GitHub releases, and download the latest/specified version
	// this will also mean we can get the size

	for _, pluginURL := range pluginURLs {
		// TODO: get replace from flag?

		// TODO: get size. Can't do that for github URLs, because
		// they're in S3, and so you can't HEAD them to find out
		// how big they are. But can do this for other URLs, so at
		// least attempt it

		err := lib.DownloadFile(pluginURL, pluginDir, 0, false)
		if err != nil {
			return fmt.Errorf("Could not download from URL: %s", err)
		}

		// chmod +x the file (well, -rwxr-xr-x, but close enough)
		filename := pluginDir + "/" + path.Base(pluginURL)
		os.Chmod(filename, 0755)
		if err != nil {
			return fmt.Errorf("Could not make plugin executable: %s", err)
		}

	}

	return nil
}

// fixTerraformImageForCustomPlugins applies a workaround for
// https://github.com/hashicorp/docker-hub-images/pull/63
// to a specified docker image
//
// If possible, as a workaround, run `apk add libc6-compat` in terraform container first
// e.g. with something like...
// https://github.com/edupo/cali/blob/3eddce060ececa2790fe6908b9f7441aac40fc3f/docker/container.go#L109
// run image, then overwrite local version
func fixTerraformImageForCustomPlugins(c *cali.Task, imageName string) error {
	log.Debugf("begin fixTerraformImageForCustomPlugins")

	// Always pull the image
	if err := c.PullImage(c.Conf.Image); err != nil {
		return fmt.Errorf("Failed to fetch image: %s", err)
	}

	// Create a container, with modified config
	c.Conf.Entrypoint = []string{"/sbin/apk"}
	c.Conf.Cmd = []string{"add", "--update", "libc6-compat"}
	resp, err := c.Cli.ContainerCreate(context.Background(), c.Conf, c.HostConf,
		c.NetConf, "")
	if err != nil {
		return fmt.Errorf("Failed to create container: %s", err)
	}

	// Start the container
	if err := c.Cli.ContainerStart(context.Background(), resp.ID,
		types.ContainerStartOptions{}); err != nil {
		return err
	}

	// TODO: Wait until container has finished running
	time.Sleep(10000 * time.Millisecond)

	// set container config, for saving the image
	c.Conf.Entrypoint = []string{"/bin/terraform"}
	c.Conf.Cmd = []string{}

	// Commit (save back to image name)
	if _, err := c.Cli.ContainerCommit(context.Background(), resp.ID,
		types.ContainerCommitOptions{
			Reference: imageName,
			// TODO: config
			Config: c.Conf,
		},
	); err != nil {
		return err
	}

	// Delete the container; we no longer need it
	if err = c.DeleteContainer(resp.ID); err != nil {
		return fmt.Errorf("Failed to remove container: %s", err)
	}

	log.Debugf("end fixTerraformImageForCustomPlugins")
	return nil
}
