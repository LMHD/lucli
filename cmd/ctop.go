package cmd

import "github.com/skybet/cali"

func init() {

	command := cli.NewCommand("ctop [command]")
	command.SetShort("Top for containers")

	task := command.Task("quay.io/vektorlab/ctop:latest")

	task.SetInitFunc(func(t *cali.Task, args []string) {
		_ = task.BindDockerSocket()
	})

}
