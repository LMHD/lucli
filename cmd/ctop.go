package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/skybet/cali"
)

func init() {

	command := cli.NewCommand("ctop [command]")
	command.SetShort("Top for containers")

	task := command.Task("quay.io/vektorlab/ctop:latest")

	dockerSocket, err := task.Bind("/var/run/docker.sock", "/var/run/docker.sock")
	if err != nil {
		log.Fatalf("Unable to format Docker socket bind: %s", err)
	}
	task.AddBinds([]string{dockerSocket})

	task.SetInitFunc(func(t *cali.Task, args []string) {
		// TODO CALI: This needs to exist, even if its empty
	})

}
