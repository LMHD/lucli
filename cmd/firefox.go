package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/skybet/cali"

	"github.com/lmhd/lucli/lib"
)

func init() {

	command := cli.NewCommand("firefox [params]")
	command.SetShort("Run an isolated web browser in an ephemeral container")

	task := command.Task("jess/firefox")

	// Unsure if this is even necessary
	tmpX, err := task.Bind("/tmp/.X11-unix", "/tmp/.X11-unix")
	if err != nil {
		log.Fatalf("Unable to bind tmp X11: %s", err)
	}
	task.AddBind(tmpX)

	// Need more than the default SHM
	// --shm-size=2GB
	// TODO: configurable with flag?
	task.HostConf.ShmSize = 2147483648
	// Run with host network
	// --net=host
	task.HostConf.NetworkMode = "host"

	// TODO: look into IPC settings for separation of shared memory
	// https://docs.docker.com/engine/reference/run/#ipc-settings-ipc
	// Should be an optional flag
	// Maybe also look into a Firefox plugin to make clear which browser is which, if running multiple?

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
