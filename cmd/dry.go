package cmd

import (
	log "github.com/sirupsen/logrus"
)

func init() {

	command := cli.NewCommand("dry [command]")
	command.SetShort("Manage docker")

	task := command.Task("docker.io/moncho/dry")

	dockerSocket, err := task.Bind("/var/run/docker.sock", "/var/run/docker.sock")
	if err != nil {
		log.Fatalf("Unable to format Docker socket bind: %s", err)
	}
	task.AddBinds([]string{dockerSocket})

}
