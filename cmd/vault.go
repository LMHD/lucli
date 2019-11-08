package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/skybet/cali"
)

func init() {

	command := cli.NewCommand("vault [command]")

	command.SetShort("Run Vault in an ephemeral container")

	// Set default image for Run function
	imageName := "docker.io/library/vault"
	imageVersion := "latest"
	command.Flags().String("vault-version", imageVersion, "Version of image to use")

	// Default Vault server
	vaultAddress := "https://vault.fancycorp.io:2082"
	command.Flags().String("vault-address", vaultAddress, "Vault server to talk to")

	// Bind all flags
	command.BindFlags()

	task := command.Task(fmt.Sprintf("%s:%s", imageName, imageVersion))

	vaultToken, err := task.Bind("~/.vault-token", "/root/.vault-token")
	if err != nil {
		log.Fatalf("Unable to format Vault token bind: %s", err)
	}
	task.AddBinds([]string{vaultToken})

	task.Conf.Entrypoint = []string{"/bin/vault"}

	// Init function, set profile, and image version
	task.SetInitFunc(func(t *cali.Task, args []string) {
		t.AddEnv("VAULT_ADDR", cli.FlagValues().GetString("vault-address"))
		t.SetImage(fmt.Sprintf("%s:%s", imageName, cli.FlagValues().GetString("vault-version")))
	})

}
