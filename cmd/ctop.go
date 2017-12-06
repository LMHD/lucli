package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/skybet/cali"
)

func CtopCli(cli *cali.Cli) {

	ctop := cli.NewCommand("ctop [command]")
	ctop.SetShort("Top for containers")

	ctopTask := ctop.Task("quay.io/vektorlab/ctop:latest")

	dockerSocket, err := ctopTask.Bind("/var/run/docker.sock", "/var/run/docker.sock")
	if err != nil {
		log.Fatalf("Unable to bind docker socket: %s", err)
	}
	ctopTask.AddBinds([]string{dockerSocket})

	ctopTask.SetInitFunc(func(t *cali.Task, args []string) {
		// TODO CALI: This needs to exist, even if its empty
	})

}
