package aufgabe3

import (
	"BWINF/Aufgabe3/pancake"
	"BWINF/cli"
	"flag"
)

var Aufgabe3 = cli.Command{
	Name:        "aufgabe3",
	Usage:       "aufgabe3 <subcommand>",
	Description: "command für aufgabe3",
	Flags:       aufgabe3Flags,
}

var aufgabe3Flags = flag.NewFlagSet("aufgabe3", flag.ExitOnError)

func init() {
	aufgabe3Flags.BoolVar(&pancake.KeepTrackOfSide, "keepTrackOfSide", false, "Ob am Schluss der Stack zur selben Seite zeigen soll")

	Aufgabe3.AddCommand(Sort)
	Aufgabe3.AddCommand(Pwue)
}
