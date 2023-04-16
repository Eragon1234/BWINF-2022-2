package ant

type Config struct {
	NumOfAnts                    int
	NumOfIterations              int
	RandomChance                 float64
	PheromoneWeight              float64
	DistanceWeight               float64
	PheromoneDecreasement        float64
	EliteProportion              float64
	CutoffWhenLongerThanShortest float64
}
