package ant

import "runtime"

type Config struct {
	NumOfAnts            int
	NumOfIterations      int
	PheromoneWeight      float64
	DistanceWeight       float64
	PheromoneAmount      float64
	PheromoneEvaporation float64
	Elite                int
}

var DefaultConfig = Config{
	NumOfAnts:            runtime.NumCPU() * 16,
	NumOfIterations:      10000,
	PheromoneWeight:      1,
	DistanceWeight:       3,
	PheromoneEvaporation: 0.1,
	PheromoneAmount:      1,
	Elite:                runtime.NumCPU() * 4,
}
