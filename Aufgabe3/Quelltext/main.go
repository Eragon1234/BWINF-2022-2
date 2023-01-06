package main

import (
	"Aufgabe3/pancake"
	"flag"
	"fmt"
	"log"
	"os"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "", "Name der Datei mit dem Pancake")
	flag.StringVar(&filename, "file", "", "Name der Datei mit dem Pancake")
	flag.StringVar(&filename, "filename", "", "Name der Datei mit dem Pancake")
	flag.Parse()
}

func main() {
	if filename == "" && len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		flag.PrintDefaults()
		return
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

	p, err := pancake.ParsePancakeFromReader(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", p)
}
