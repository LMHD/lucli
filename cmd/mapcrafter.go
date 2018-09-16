package cmd

import (
	"github.com/skybet/cali"
)

func init() {

	command := cli.NewCommand("mapcrafter [command]")

	command.SetShort("Run Mapcrafter in an ephemeral container")

	task := command.Task("lucymhdavies/mapcrafter")

	// Init function
	task.SetInitFunc(func(t *cali.Task, args []string) {
	})

}
