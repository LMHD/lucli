package cmd

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/lmhd/lucli/creds"
	"github.com/lmhd/lucli/lib"
	"github.com/skybet/cali"
)

func init() {

	// TODO: alias as awsshell?
	command := cli.NewCommand("awscli [command]")
	command.SetShort("AWS CLI / Shell")

	// Allow this command to be run with awsshell too
	// Dunno. Might made this a top tier command instead, so it shows in help.
	command.SetAliases([]string{"awsshell"})

	// TODO: long description
	// Include AWS CLI/Shell
	// Mention JQ is included

	command.Flags().BoolP("shell", "s", false, "Use AWS Shell")
	command.BindFlags()

	task := command.Task("lucymhdavies/awscli:latest")

	task.SetInitFunc(func(t *cali.Task, args []string) {

		// if awscli -s, or awsshell
		if cli.FlagValues().GetBool("shell") || lib.StringSliceContains(os.Args, "awsshell") {
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
