package aufgabe1

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/cli"
	"errors"
	"fmt"
	"os"
	"runtime"
)

var Ant = cli.Command{
	Name:        "ant",
	Usage:       "ant <filename>",
	Description: "berechnet eine mögliche Lösung mit einer Ameisenkoloniesimulation",
	Action:      ant,
}

var cfg = graph.Config{
	NumOfAnts:                    runtime.NumCPU() * 128,
	NumOfIterations:              100_000,
	RandomChance:                 0.6,
	PheromoneWeight:              0.5,
	DistanceWeight:               0.7,
	PheromoneDecreasement:        0.8,
	EliteProportion:              0.1,
	CutoffWhenLongerThanShortest: 3.0,
}

func ant(args []string, cmd *cli.Command) error {
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
	path := graph.VisitAllAntColonyOptimization(cfg, g)
	fmt.Println("Length of path:", graph.LengthOfPheromonePath(path))
	for _, edge := range path {
		fmt.Printf("%v -> %v\n", edge.From.Name, edge.To.Name)
	}
	return nil
}
