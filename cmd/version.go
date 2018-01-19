package cmd

import (
	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/skybet/cali"

	"github.com/lmhd/lucli/lib"
)

func init() {

	command := cli.NewCommand("version")
	command.SetShort("Which version are we running?")

	var taskFunc cali.TaskFunc = func(t *cali.Task, args []string) {
		log.Infof("GOOS: %s, GOARCH: %s", runtime.GOOS, runtime.GOARCH)
		log.Infof("lucli v%s (%s) (%s)", lib.Version, lib.BuildTime, lib.BuildCommit)
	}

	command.Task(taskFunc)

}
