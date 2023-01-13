package main

import (
	"Aufgabe3/pancake"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "", "Pfad zur Datei mit dem PancakeStack")
	flag.StringVar(&filename, "file", "", "Pfad zur Datei mit dem PancakeStack")
	flag.StringVar(&filename, "filename", "", "Pfad zur Datei mit dem PancakeStack")
	flag.Parse()
}

func main() {
	if filename == "" && len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		flag.PrintDefaults()
		return
	}

	if !fs.ValidPath(filename) {
		log.Fatal("Invalid filepath")
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	p, err := pancake.Parse(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("PancakeStack Sort Way: %v", pancake.ShortestBruteForceSortSteps(p))
}
