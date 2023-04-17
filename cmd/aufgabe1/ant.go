package aufgabe1

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/Aufgabe1/graph/ant"
	"BWINF/cli"
	"errors"
	"flag"
	"fmt"
	"os"
)

var AntFlags = flag.NewFlagSet("ant", flag.ExitOnError)

var Ant = cli.Command{
	Name:        "ant",
	Usage:       "ant <filename>",
	Flags:       AntFlags,
	Description: "berechnet eine mögliche Lösung mit einer Ameisenkoloniesimulation",
	Action:      antCommand,
}

func init() {
	AntFlags.IntVar(&ant.DefaultConfig.NumOfAnts, "numOfAnts", ant.DefaultConfig.NumOfAnts, "Anzahl der Ameisen")
	AntFlags.IntVar(&ant.DefaultConfig.NumOfIterations, "numOfIterations", ant.DefaultConfig.NumOfIterations, "Anzahl der Iterationen")
	AntFlags.Float64Var(&ant.DefaultConfig.PheromoneWeight, "pheromoneWeight", ant.DefaultConfig.PheromoneWeight, "Gewichtung der Pheromone")
	AntFlags.Float64Var(&ant.DefaultConfig.DistanceWeight, "distanceWeight", ant.DefaultConfig.DistanceWeight, "Gewichtung der Entfernung")
	AntFlags.Float64Var(&ant.DefaultConfig.PheromoneAmount, "pheromoneAmount", ant.DefaultConfig.PheromoneAmount, "Menge der Pheromone")
	AntFlags.Float64Var(&ant.DefaultConfig.PheromoneEvaporation, "pheromoneEvaporation", ant.DefaultConfig.PheromoneEvaporation, "Verdunstung der Pheromone")
	AntFlags.IntVar(&ant.DefaultConfig.Elite, "elite", ant.DefaultConfig.Elite, "Anzahl der besten Ameisen")
	AntFlags.IntVar(&ant.DefaultConfig.Patience, "patience", ant.DefaultConfig.Patience, "Anzahl der Iterationen ohne Verbesserung")
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
	path := ant.VisitAll(ant.DefaultConfig, g)
	fmt.Println("Length of path:", ant.LengthOfPheromonePath(path))
	displayPath(path)
	return nil
}
