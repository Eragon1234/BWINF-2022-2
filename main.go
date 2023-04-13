package main

import (
	"BWINF/cli"
	"BWINF/cmd/aufgabe1"
	"BWINF/cmd/aufgabe3"
	"log"
	"os"
)

var command = cli.Command{
	Name:        "BWINF",
	Usage:       "BWINF <command> [args]",
	Description: "BWINF",
}

func init() {
	command.AddCommand(aufgabe1.Aufgabe1)
	command.AddCommand(aufgabe3.Aufgabe3)
}

func main() {
	err := command.Run(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
}
