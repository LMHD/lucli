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
		log.Infof("lucli v%s (%s) (%s)", lib.Version, lib.BuildTime, lib.BuildCommit)
		log.Debugf("GOOS: %s, GOARCH: %s", runtime.GOOS, runtime.GOARCH)
	}

	command.Task(taskFunc)

}
