package aufgabe1

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/cli"
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
	NumOfAnts:                    runtime.NumCPU() * 64,
	NumOfIterations:              10_000,
	RandomChance:                 0.5,
	PheromoneWeight:              0.3,
	DistanceWeight:               0.7,
	PheromoneDecreasement:        0.7,
	EliteProportion:              0.1,
	CutoffWhenLongerThanShortest: 5.0,
}

func ant(args []string, cmd *cli.Command) error {
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
		fmt.Printf("%v -> %v\n", edge.From, edge.To)
	}
	return nil
}
