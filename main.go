package main

import (
	"github.com/skybet/cali"

	"github.com/lmhd/lucli/cmd"
)

func main() {
	cli := cali.NewCli("lucli")
	cli.SetShort("Example CLI tool")
	cli.SetLong("A nice long description of what your tool actually does")

	cmd.TerraformCli(cli)
	cmd.FrazelleCli(cli)

	cli.Start()
}
