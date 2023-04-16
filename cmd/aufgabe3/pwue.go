package aufgabe3

import (
	"BWINF/Aufgabe3/pancake/sort"
	"BWINF/cli"
	"errors"
	"fmt"
	"strconv"
)

var Pwue = cli.Command{
	Name:        "pwue",
	Usage:       "pwue <n>",
	Description: "Berechnet die PWUE Zahl für n",
	Action:      pwue,
}

func pwue(args []string, _ *cli.Command) error {
	if len(args) == 0 {
		return errors.New("missing n")
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("n muss eine Zahl sein")
	}
	stack, sortSteps := sort.CalculatePWUE(n)
	fmt.Printf("PWUE(%v) = %v\n", n, len(sortSteps))
	fmt.Println("Beispiel Stack:", stack)
	fmt.Printf("Sortierungs Schritte: %v\n", sortSteps)
	return nil
}
