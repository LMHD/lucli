package cmd

import (
	"github.com/lmhd/lucli/creds"
	log "github.com/sirupsen/logrus"
	"github.com/skybet/cali"
)

func init() {

	// TODO: alias as awsshell?
	command := cli.NewCommand("awscli [command]")
	command.SetShort("AWS CLI / Shell")

	// TODO: long description
	// Include AWS CLI/Shell
	// Mention JQ is included

	command.Flags().BoolP("shell", "s", false, "Use AWS Shell")
	command.BindFlags()

	task := command.Task("docker.io/lucymhdavies/awscli:latest")

	task.SetInitFunc(func(t *cali.Task, args []string) {
		// TODO: detect if ran with awsshell alias?

		if cli.FlagValues().GetBool("shell") {
			log.Debugf("Using AWS Shell")

			// TODO: cali DockerClient.SetEntrypoint
			t.Conf.Entrypoint = []string{"/usr/local/bin/aws-shell"}
		} else {
			log.Debugf("Using AWS CLI")
		}

		task.AddEnv("AWS_PROFILE", cli.FlagValues().GetString("aws-profile"))
		_ = creds.BindAWS(t, args)
	})

}
