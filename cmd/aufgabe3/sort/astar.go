package sort

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/Aufgabe3/pancake/sort"
	"BWINF/cli"
	"errors"
	"io/fs"
	"os"
)

var Astar = cli.Command{
	Name:        "astar",
	Usage:       "astar <filename>",
	Description: "Berechnet den Sortierweg für den Stack in der Datei mit A*",
	Action:      sortAstar,
}

func sortAstar(args []string, _ *cli.Command) error {
	if len(args) == 0 {
		return errors.New("keine Datei angegeben")
	}
	filename := args[0]
	if !fs.ValidPath(filename) {
		return errors.New("ungültiger Dateipfad")
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	p, err := pancake.ParseStack(file)
	if err != nil {
		return err
	}
	printStack(p)
	sortSteps := sort.Astar(p)
	printSortSteps(p, sortSteps)
	return nil
}
