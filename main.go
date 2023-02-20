package main

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/Aufgabe3/pancake"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "", "Pfad zur Eingabedatei")
	flag.StringVar(&filename, "file", "", "Pfad zur Eingabedatei")
	flag.StringVar(&filename, "filename", "", "Pfad zur Eingabedatei")
	flag.Parse()
}

func main() {
	command := flag.Arg(0)
	if command == "" {
		fmt.Println("Missing command")
		flag.Usage()
		return
	}
	if filename == "" {
		filename = flag.Arg(1)
	}
	if filename == "" {
		flag.Usage()
		return
	}

	if !fs.ValidPath(filename) {
		log.Println("Invalid filepath")
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
	case "help":
		flag.Usage()
		return
	case "aufgabe1":
		weightedGraph, err := graph.ParseComplete(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(weightedGraph)
	case "aufgabe3":
		p, err := pancake.ParseStack[uint8](file)
		if err != nil {
			log.Println("Failed to parse pancake")
			return
		}

		fmt.Println("Pancake:", p)
		for _, step := range pancake.BruteForceSort(p) {
			fmt.Printf("Flip at %vth, new pancake %v\n", step, *p.Flip(int(step)))
		}
	default:
		fmt.Println("Unknown command or missing command")
		flag.Usage()
		return
	}
}
