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
	NumOfAnts:            runtime.NumCPU() * 256,
	NumOfIterations:      300,
	PheromoneWeight:      1,
	DistanceWeight:       3,
	PheromoneEvaporation: 0.03,
	PheromoneAmount:      1,
	Elite:                runtime.NumCPU() * 256,
}
