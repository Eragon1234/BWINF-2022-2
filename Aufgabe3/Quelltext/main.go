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
		log.Fatalln("Invalid filepath")
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

	p, err := pancake.Parse[uint8](file)
	if err != nil {
		log.Fatalln("Failed to parse pancake")
	}

	fmt.Println("Pancake:", p)
	for _, step := range pancake.BruteForceSort(p) {
		fmt.Printf("Flip at %vth, new pancake %v\n", step, *p.Flip(int(step)))
	}
}
