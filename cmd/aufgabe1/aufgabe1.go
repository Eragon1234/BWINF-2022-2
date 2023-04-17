package aufgabe1

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/cli"
	"fmt"
)

var Aufgabe1 = cli.Command{
	Name:        "aufgabe1",
	Usage:       "aufgabe1 <subcommand>",
	Description: "command für aufgabe1",
	Action:      antCommand,
	Subcommands: []cli.Command{
		Ant,
		Shortest,
	},
}

func displayPath[T any, M any](path []graph.Edge[T, M]) {
	for _, edge := range path {
		fmt.Printf("%v -> %v\n", edge.From.Name, edge.To.Name)
	}
}
