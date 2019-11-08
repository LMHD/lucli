package cmd

import (
	"github.com/lmhd/lucli/creds"
	"github.com/skybet/cali"
)

func init() {

	command := cli.NewCommand("s3explorer")
	command.SetShort("S3 Explorer")

	command.BindFlags()

	task := command.Task("docker.io/lucymhdavies/s3explorer:latest")

	task.SetInitFunc(func(t *cali.Task, args []string) {
		task.AddEnv("AWS_PROFILE", cli.FlagValues().GetString("aws-profile"))
		_ = creds.BindAWS(t, args)
	})

}
