package optimize

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/Aufgabe1/graph/ant"
	"BWINF/Aufgabe1/vector"
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func Optimize(g graph.WeightedGraph[vector.Coordinate, graph.DistanceAngle]) ant.Config {
	var bestConfig ant.Config
	var bestLength float64
	var run = true
	go terminator(&run)
	var batchSize = 2
	var wg sync.WaitGroup
	for run {
		fmt.Fprintf(os.Stderr, "best length %v, %v\n", bestLength, bestConfig)
		wg.Add(batchSize)
		for i := 0; i < batchSize; i++ {
			go func() {
				defer wg.Done()
				cfg := RandomConfig()
				if cfg.Elite > cfg.NumOfAnts {
					cfg.Elite = cfg.NumOfAnts
				}
				path := ant.VisitAllAntColonyOptimization(cfg, g)
				length := ant.LengthOfPheromonePath(path)
				if bestLength == 0 || length < bestLength {
					bestLength = length
					bestConfig = cfg
				}
			}()
		}
		wg.Wait()

		print("\033[F\033[2K")
	}
	return bestConfig
}

func terminator(run *bool) {
	fmt.Scanln()
	*run = false
}

func RandomConfig() ant.Config {
	return ant.Config{
		NumOfAnts:            rand.Intn(2000),
		NumOfIterations:      rand.Intn(3000),
		PheromoneWeight:      float64(rand.Intn(5)),
		DistanceWeight:       float64(rand.Intn(5)),
		PheromoneAmount:      float64(rand.Intn(5)),
		PheromoneEvaporation: rand.Float64() / 2,
		Elite:                rand.Intn(3000),
	}
}
