package main

import (
	"Aufgabe1/graph"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "", "Pfad zur Datei mit den Koordinaten")
	flag.StringVar(&filename, "file", "", "Pfad zur Datei mit den Koordinaten")
	flag.StringVar(&filename, "filename", "", "Pfad zur Datei mit den Koordinaten")
	flag.Parse()
}

func main() {
	if filename == "" && len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
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

	weightedGraph, err := graph.ParseComplete(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(weightedGraph)
}
