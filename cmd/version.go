package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/skybet/cali"

	"github.com/lmhd/lucli/lib"
)

func init() {

	command := cli.NewCommand("version")
	command.SetShort("Which version are we running?")

	command.Flags().Bool("check-update", true, "When displaying current version, skip checking for update")
	command.BindFlags()

	var taskFunc cali.TaskFunc = func(t *cali.Task, args []string) {
		lib.PrintVersion()

		if cli.FlagValues().GetBool("check-update") {

			isLatestVersion, releaseData, err := lib.IsLatestVersion()
			if err != nil {
				log.Fatalf("Unable to check for update: %s", err)
			}

			if isLatestVersion {
				log.Infof("You're already using the latest version. Well done!")
			} else {
				log.Infof("You're not running the latest version 😱")
				log.Infof("Update to v%s with: lucli update", releaseData.Name)
			}

		}
	}

	// Simple task, just runs a function
	command.Task(taskFunc)

}
