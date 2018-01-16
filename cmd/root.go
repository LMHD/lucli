package cmd

import (
	"github.com/skybet/cali"
)

var (
	// Define this here, then all other files in cmd can add subcommands to it
	cli = cali.NewCli("lucli")
)

func init() {
	cli.SetShort("Example CLI tool")
	cli.SetLong("A nice long description of what your tool actually does")
}

func Execute() {
	cli.Start()
}
