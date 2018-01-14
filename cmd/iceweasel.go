package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/skybet/cali"

	"github.com/lmhd/lucli/lib"
)

func IceweaselCli(cli *cali.Cli) {

	command := cli.NewCommand("iceweasel [params]")
	command.SetShort("Run Iceweasel in an ephemeral container")

	task := command.Task("jess/iceweasel")

	// Unsure if this is even necessary
	tmpX, err := task.Bind("/tmp/.X11-unix", "/tmp/.X11-unix")
	if err != nil {
		log.Fatalf("Unable to bind tmp X11: %s", err)
	}
	task.AddBind(tmpX)

	task.SetInitFunc(func(t *cali.Task, args []string) {
		// Possibly move this call into lib.GetDisplay?
		err := lib.StartXQuartz()
		if err != nil {
			log.Fatalf("Unable to start X11 :(")
		}

		display, err := lib.GetDisplay()
		if err != nil {
			log.Fatalf("Unable to get DISPLAY :(")
		}

		t.AddEnv("DISPLAY", display)
	})

}
