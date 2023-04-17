package ant

import "runtime"

// Config is the configuration for the ant algorithm.
type Config struct {
	NumOfAnts            int     // NumOfAnts is the number of ant workers to run per iteration.
	NumOfIterations      int     // NumOfIterations is the number of iterations to run.
	PheromoneWeight      float64 // PheromoneWeight is the weight of the pheromones when choosing the next node.
	DistanceWeight       float64 // DistanceWeight is the weight of the distance when choosing the next node.
	PheromoneAmount      float64 // PheromoneAmount is the amount of pheromones to add to the graph.
	PheromoneEvaporation float64 // PheromoneEvaporation is a factor to reduce the pheromones by.
	Elite                int     // Elite is the number of best ants to give pheromones.
	Patience             int     // Patience is the number of iterations without improvement before the algorithm stops.
}

// DefaultConfig is a default configuration that should give decent results.
var DefaultConfig = Config{
	NumOfAnts:            runtime.NumCPU() * 512,
	NumOfIterations:      100,
	PheromoneWeight:      2,
	DistanceWeight:       6,
	PheromoneEvaporation: 0.1,
	PheromoneAmount:      1,
	Elite:                runtime.NumCPU(),
	Patience:             25,
}
