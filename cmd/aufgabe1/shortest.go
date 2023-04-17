package aufgabe1

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/cli"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

var Shortest = cli.Command{
	Name:        "shortest",
	Usage:       "shortest <filename>",
	Description: "berechnet eine Strecke die alle Koordinaten besucht basierend auf der kürzesten Strecke",
	Action:      shortest,
}

func shortest(args []string, c *cli.Command) error {
	if len(args) == 0 {
		return errors.New("missing filename")
	}
	filename := args[0]
	if !fs.ValidPath(filename) {
		return errors.New("invalid filename")
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	g, err := graph.ParseComplete(file)
	if err != nil {
		return err
	}
	edges := graph.VisitAllShortestEdge(g)
	fmt.Println("Länge der Strecke:", graph.LengthOfPath(edges))
	displayPath(edges)
	return nil
}
