package aufgabe3

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/Aufgabe3/pancake/sort"
	"BWINF/cli"
	sortCmd "BWINF/cmd/aufgabe3/sort"
	"flag"
	"runtime"
)

var Aufgabe3 = cli.Command{
	Name:        "aufgabe3",
	Usage:       "aufgabe3 <subcommand>",
	Description: "command für aufgabe3",
	Flags:       aufgabe3Flags,
	Subcommands: []cli.Command{
		sortCmd.Sort,
		Pwue,
	},
}

var aufgabe3Flags = flag.NewFlagSet("aufgabe3", flag.ExitOnError)

func init() {
	aufgabe3Flags.BoolVar(&pancake.KeepTrackOfSide, "keepTrackOfSide", false, "Ob am Schluss der Stack zur selben Seite zeigen soll")
	aufgabe3Flags.IntVar(&sort.WorkerCount, "workerCount", runtime.NumCPU(), "Anzahl der Worker die gleichzeitig arbeiten sollen")
}
