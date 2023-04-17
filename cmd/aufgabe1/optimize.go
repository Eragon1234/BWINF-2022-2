package aufgabe1

import (
	"BWINF/Aufgabe1/graph"
	optimize2 "BWINF/Aufgabe1/optimize"
	"BWINF/cli"
	"fmt"
	"os"
)

var Optimize = cli.Command{
	Name:        "optimize",
	Usage:       "optimize <filename>",
	Description: "optimiert die Parameter für die Ameisen",
	Action:      optimize,
}

func optimize(args []string, c *cli.Command) error {
	filename := args[0]
	file, _ := os.Open(filename)
	defer file.Close()
	g, _ := graph.ParseComplete(file)

	fmt.Println(optimize2.Optimize(g))
	return nil
}
