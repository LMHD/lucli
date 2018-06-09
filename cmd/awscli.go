package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/lmhd/lucli/creds"
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

	task := command.Task("lucymhdavies/awscli:latest")

	task.SetInitFunc(func(t *cali.Task, args []string) {
		// TODO: detect if ran with awsshell alias?

		if cli.FlagValues().GetBool("shell") {
			log.Debugf("Using AWS Shell")

			// TODO: cali DockerClient.SetEntrypoint
			t.Conf.Entrypoint = []string{"/usr/local/bin/aws-shell"}
		} else {
			log.Debugf("Using AWS CLI")
		}

		_ = creds.BindAWS(t, args)
	})

}
