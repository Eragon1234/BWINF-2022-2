package aufgabe1

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/Aufgabe1/graph/ant"
	"BWINF/cli"
	"errors"
	"fmt"
	"os"
)

var Ant = cli.Command{
	Name:        "ant",
	Usage:       "ant <filename>",
	Description: "berechnet eine mögliche Lösung mit einer Ameisenkoloniesimulation",
	Action:      antCommand,
}

func antCommand(args []string, cmd *cli.Command) error {
	if len(args) == 0 {
		return errors.New("missing filename")
	}
	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	g, err := graph.ParseComplete(file)
	if err != nil {
		return err
	}
	path := ant.VisitAllAntColonyOptimization(ant.DefaultConfig, g)
	fmt.Println("Length of path:", ant.LengthOfPheromonePath(path))
	for _, edge := range path {
		fmt.Printf("%v -> %v\n", edge.From.Name, edge.To.Name)
	}
	return nil
}
