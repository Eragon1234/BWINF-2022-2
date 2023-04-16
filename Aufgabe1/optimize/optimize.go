package optimize

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/Aufgabe1/graph/ant"
	"BWINF/Aufgabe1/vector"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"sync"
)

func OptimizeParameters(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()
	g, _ := graph.ParseComplete(file)
	population := make([]Individual, 10)

	for i := range population {
		population[i].RandomizeParameters()
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		fmt.Fprintln(os.Stderr, "Generation:", i)
		fmt.Fprintln(os.Stderr, "Best:", population[0].ResultLength, population[0])
		wg.Add(len(population))
		for j := range population {
			go func(j int) {
				defer wg.Done()
				population[j].Run(g.Copy())
			}(j)
		}
		wg.Wait()
		sort.Slice(population, func(i, j int) bool {
			return population[i].ResultLength < population[j].ResultLength
		})
		for j := range population {
			if j < len(population)/2 {
				continue
			}
			switch rand.Intn(4) {
			case 0:
				population[j].RandomizeRandomParameter()
			case 1, 2:
				population[j].TakeTwo(population[rand.Intn(len(population)/2)])
			case 3:
				population[j].TakeAndModify(population[rand.Intn(len(population)/2)])
			}
		}

		fmt.Fprint(os.Stderr, "\u001B[1A\u001B[K")
		fmt.Fprint(os.Stderr, "\u001B[1A\u001B[K")
	}

	for i := range population {
		fmt.Printf("%v\n", population[i])
	}
}

type Individual struct {
	NumOfAnts                    int
	NumOfIterations              int
	RandomChance                 float64
	PheromoneWeight              float64
	DistanceWeight               float64
	PheromoneDecreasement        float64
	EliteProportion              float64
	CutoffWhenLongerThanShortest float64
	ResultLength                 float64
}

func (i *Individual) Run(weightedGraph graph.WeightedGraph[vector.Coordinate, graph.DistanceAngle]) {
	config := ant.Config{
		NumOfAnts:                    i.NumOfAnts,
		NumOfIterations:              i.NumOfIterations,
		RandomChance:                 i.RandomChance,
		PheromoneWeight:              i.PheromoneWeight,
		DistanceWeight:               i.DistanceWeight,
		PheromoneDecreasement:        i.PheromoneDecreasement,
		EliteProportion:              i.EliteProportion,
		CutoffWhenLongerThanShortest: i.CutoffWhenLongerThanShortest,
	}
	i.ResultLength = ant.LengthOfPheromonePath(ant.VisitAllAntColonyOptimization(config, weightedGraph))
}

func (i *Individual) RandomizeParameters() {
	i.NumOfAnts = rand.Intn(runtime.NumCPU()*128) + 1
	i.NumOfIterations = rand.Intn(3000)
	i.RandomChance = rand.Float64()
	i.PheromoneWeight = rand.Float64()
	i.DistanceWeight = rand.Float64()
	i.PheromoneDecreasement = rand.Float64()
	i.EliteProportion = rand.Float64()
	i.CutoffWhenLongerThanShortest = rand.Float64() * 5
}

func (i *Individual) RandomizeRandomParameter() {
	switch rand.Intn(8) {
	case 0:
		i.NumOfAnts = rand.Intn(runtime.NumCPU()*128) + 1
	case 1:
		i.NumOfIterations = rand.Intn(3000)
	case 2:
		i.RandomChance = rand.Float64()
	case 3:
		i.PheromoneWeight = rand.Float64()
	case 4:
		i.DistanceWeight = rand.Float64()
	case 5:
		i.PheromoneDecreasement = rand.Float64()
	case 6:
		i.EliteProportion = rand.Float64()
	case 7:
		i.CutoffWhenLongerThanShortest = rand.Float64() * 5
	}
}

func (i *Individual) TakeTwo(i2 Individual) {
	switch rand.Intn(4) {
	case 0:
		i.NumOfAnts = i2.NumOfAnts
		i.DistanceWeight = i2.DistanceWeight
	case 1:
		i.NumOfIterations = i2.NumOfIterations
		i.PheromoneWeight = i2.PheromoneWeight
	case 2:
		i.RandomChance = i2.RandomChance
		i.PheromoneDecreasement = i2.PheromoneDecreasement
	case 3:
		i.EliteProportion = i2.EliteProportion
		i.CutoffWhenLongerThanShortest = i2.CutoffWhenLongerThanShortest
	}
}

func (i *Individual) ModifyRandomParameter() {
	sign := rand.Intn(2)
	if sign == 0 {
		sign = -1
	}
	switch rand.Intn(8) {
	case 0:
		i.NumOfAnts += rand.Intn(100) * sign
		if i.NumOfAnts < 1 {
			i.NumOfAnts = 1
		}
	case 1:
		i.NumOfIterations += rand.Intn(100) * sign
		if i.NumOfIterations < 1 {
			i.NumOfIterations = 1
		}
	case 2:
		i.RandomChance += rand.Float64() / 10 * float64(sign)
	case 3:
		i.PheromoneWeight += rand.Float64() / 10 * float64(sign)
	case 4:
		i.DistanceWeight += rand.Float64() / 10 * float64(sign)
	case 5:
		i.PheromoneDecreasement += rand.Float64() / 10 * float64(sign)
	case 6:
		i.EliteProportion += rand.Float64() / 10 * float64(sign)
	case 7:
		i.CutoffWhenLongerThanShortest += rand.Float64() * float64(sign)
		if i.CutoffWhenLongerThanShortest < 1.5 {
			i.CutoffWhenLongerThanShortest = 1.5
		}
	}
}

func (i *Individual) TakeAndModify(i2 Individual) {
	i.NumOfAnts = i2.NumOfAnts
	i.NumOfIterations = i2.NumOfIterations
	i.RandomChance = i2.RandomChance
	i.PheromoneWeight = i2.PheromoneWeight
	i.DistanceWeight = i2.DistanceWeight
	i.PheromoneDecreasement = i2.PheromoneDecreasement
	i.EliteProportion = i2.EliteProportion
	i.CutoffWhenLongerThanShortest = i2.CutoffWhenLongerThanShortest

	i.ModifyRandomParameter()
	i.ModifyRandomParameter()
}
