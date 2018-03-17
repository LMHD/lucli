package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/skybet/cali"

	"github.com/lmhd/lucli/lib"
)

func init() {

	command := cli.NewCommand("version")
	command.SetShort("Which version are we running?")

	var taskFunc cali.TaskFunc = func(t *cali.Task, args []string) {
		log.Infof("lucli v%s", lib.Version)
		log.Debugf("Build Date: %s", lib.BuildTime)
		log.Debugf("Git Details: %s @ %s", lib.BuildCommit, lib.BuildRepo)

		// TODO: check for updates
	}

	// Simple task, just runs a function
	command.Task(taskFunc)

}
