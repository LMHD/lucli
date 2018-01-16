package cmd

import (
	log "github.com/Sirupsen/logrus"
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

}
