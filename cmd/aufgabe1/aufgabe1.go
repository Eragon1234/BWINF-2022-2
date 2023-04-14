package aufgabe1

import "BWINF/cli"

var Aufgabe1 = cli.Command{
	Name:        "aufgabe1",
	Usage:       "aufgabe1 <subcommand>",
	Description: "command für aufgabe1",
	Action:      ant,
}

func init() {
	Aufgabe1.AddCommand(Ant)
	Aufgabe1.AddCommand(Shortest)
}
