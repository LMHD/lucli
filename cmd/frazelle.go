package cmd

import "github.com/skybet/cali"

func FrazelleCli(cli *cali.Cli) {

	fraz := cli.NewCommand("frazelle [command]")
	fraz.SetShort("Wanna be Jessie Frazelle?")

	//	_ = fraz.Task(func(t *Task, args []string) {
	//		fraz.Usage
	//	})

}
