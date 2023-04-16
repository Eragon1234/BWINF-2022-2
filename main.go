package main

import (
	"BWINF/cli"
	"BWINF/cmd/aufgabe1"
	"BWINF/cmd/aufgabe3"
	"fmt"
	"os"
)

var command = cli.Command{
	Name:        "BWINF",
	Usage:       "BWINF <command> [args]",
	Description: "BWINF",
	Subcommands: []cli.Command{
		aufgabe1.Aufgabe1,
		aufgabe3.Aufgabe3,
	},
}

func main() {
	err := command.Run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
