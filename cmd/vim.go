package cmd

import (
	"os/user"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command := cli.NewCommand("vim [command]")
	command.SetShort("Run Vim in an ephemeral container")
	task := command.Task("lucymhdavies/vim:latest")

	// TODO: this should be common to all containers!
	// And probably should be standardised as part of Cali!
	u, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to find uid for user: %s", err)
	}
	task.AddEnv("HOST_USER_ID", u.Uid)
	task.AddEnv("HOST_GROUP_ID", u.Gid)
}
