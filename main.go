package main

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/Aufgabe1/optimize"
	"BWINF/Aufgabe3/pancake"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

var filename string
var keepTrackOfSide bool

func init() {
	flag.StringVar(&filename, "f", "", "Pfad zur Eingabedatei")
	flag.StringVar(&filename, "file", "", "Pfad zur Eingabedatei")
	flag.StringVar(&filename, "filename", "", "Pfad zur Eingabedatei")
	flag.BoolVar(&keepTrackOfSide, "keepTrackOfSide", false, "If the pancake should keep track of the side it is on")
	flag.Parse()

	flag.Usage = func() {
		fmt.Println("Usage: <command> <filename>")
		fmt.Println()
		fmt.Println("Available commands:")
		fmt.Println("  aufgabe1 <filename>")
		fmt.Println("  aufgabe3 <filename>")
		fmt.Println()
		flag.PrintDefaults()
	}
}

func main() {
	command := flag.Arg(0)

	if command == "" {
		fmt.Println("Missing command")
		fmt.Println()
		flag.Usage()
		return
	}

	if filename == "" {
		filename = flag.Arg(1)
	}

	if !fs.ValidPath(filename) {
		fmt.Println("Invalid or missing filepath")
		fmt.Println()
		flag.Usage()
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		panic("Failed to open file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Failed to close file")
		}
	}(file)

	switch command {
	case "aufgabe1":
		weightedGraph, err := graph.ParseComplete(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		steps := graph.VisitAllShortestEdge(weightedGraph)
		//stepsAntColony := graph.VisitAllAntColonyOptimization(weightedGraph)
		fmt.Println("Steps:", graph.LengthOfPath(steps))
		//fmt.Println("Steps Ant Colony:", graph.LengthOfPheromonePath(stepsAntColony))
	case "aufgabe3":
		pancake.KeepTrackOfSide = keepTrackOfSide
		p, err := pancake.ParseStack[int8](file)
		if err != nil {
			fmt.Println("Failed to parse pancake")
			return
		}

		fmt.Println("Pancake:", p)
		for _, step := range pancake.BruteForceSort(p) {
			fmt.Printf("Flip at %vth, new pancake %v\n", step, *p.Flip(int(step)))
		}
	case "optimize":
		optimize.OptimizeParameters(filename)
	case "pwue":
		pancake.KeepTrackOfSide = keepTrackOfSide
		n, _ := strconv.Atoi(flag.Arg(2))
		stack, sortSteps := pancake.CalculatePWUE(n)
		fmt.Println("Maximum number of steps needed to sort pancake:", len(sortSteps))
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("Pancake:", stack)
		for _, step := range sortSteps {
			fmt.Printf("Flip at %vth, new pancake %v\n", step, *stack.Flip(int(step)))
		}
	default:
		fmt.Println("Unknown command or missing command")
		flag.Usage()
		return
	}
}
