package cmd

func init() {
	command := cli.NewCommand("vim [command]")
	command.SetShort("Run Vim in an ephemeral container")
	command.Task("lucymhdavies/vim:latest")
}
