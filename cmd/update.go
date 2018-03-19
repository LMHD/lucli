package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/skybet/cali"

	"github.com/lmhd/lucli/lib"
)

func init() {

	command := cli.NewCommand("update")
	command.SetShort("Update the current running version of lucli")

	var taskFunc cali.TaskFunc = func(t *cali.Task, args []string) {
		lib.PrintVersion()

		isLatestVersion, updateReleaseData, err := lib.IsLatestVersion()
		if err != nil {
			log.Fatalf("Unable to check for update: %s", err)
		}

		if isLatestVersion {
			log.Infof("You're already using the latest version. Well done!")
		} else {
			err = lib.Update(updateReleaseData)
			if err != nil {
				log.Fatalf("Unable to update: %s", err)
			}
		}
	}

	command.Task(taskFunc)

}
