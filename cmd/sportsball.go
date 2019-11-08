package cmd

func init() {

	command := cli.NewCommand("sportsball")
	command.SetShort("Sportsball! Sportsball! Sportsball!")
	command.SetLong(`
lucli sportsball

Enjoy all your favourite Sportsball highlights, powered by
https://github.com/cedricblondeau/world-cup-2018-cli-dashboard
`)

	task := command.Task("cdocker.io/edricbl/world-cup-2018-cli-dashboard")

	task.AddEnv("WITH_EMOJIS", "1")
	task.AddEnv("TZ", "Europe/London")

}
