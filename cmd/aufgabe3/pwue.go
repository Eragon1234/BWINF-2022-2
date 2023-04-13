package aufgabe3

import (
	"BWINF/Aufgabe3/pancake"
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

func pwue(args []string, cmd *cli.Command) error {
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("n muss eine Zahl sein")
	}
	stack, sortSteps := pancake.CalculatePWUE(n)
	fmt.Println("PWUE Zahl für n =", n)
	fmt.Println("Beispiel:")
	fmt.Println(stack)
	fmt.Printf("Sortierungs Schritte: %v\n", sortSteps)
	return nil
}
