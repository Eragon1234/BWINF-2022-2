package cli

import (
	"BWINF/pkg/set"
	"BWINF/pkg/slice"
	"flag"
	"fmt"
	"sort"
)

type Command struct {
	Subcommands []Command
	Flags       *flag.FlagSet
	Name        string
	Aliases     set.Set[string]
	Usage       string
	Description string
	Action      func(args []string, cmd *Command) error
}

func (c *Command) AddCommand(cmd Command) {
	c.Subcommands = append(c.Subcommands, cmd)
}

func (c *Command) Help() {
	fmt.Println(c.Usage)
	fmt.Println(c.Description)
	if c.Flags != nil {
		fmt.Println("Flags:")
		c.Flags.PrintDefaults()
	}
	if len(c.Subcommands) > 0 {
		fmt.Println("Subcommands:")
		for _, command := range c.Subcommands {
			fmt.Printf("  %v\n", command.Name)
		}
	}
}

func (c *Command) Run(args []string) error {
	sort.SliceStable(args, func(i, j int) bool {
		return args[i][0] == '-'
	})
	if c.Flags != nil {
		err := c.Flags.Parse(args)
		if err != nil {
			return err
		}
		args = c.Flags.Args()
	}
	// ignore flags
	commands := slice.FilterFunc(args, func(arg string) bool {
		return arg[0] == '-'
	})
	if len(commands) > 0 {
		if commands[0] == "help" {
			c.Help()
			return nil
		}
		for _, command := range c.Subcommands {
			if command.Name == commands[0] || command.Aliases.Contains(commands[0]) {
				return command.Run(slice.FilterFunc(args, func(arg string) bool {
					return arg == commands[0]
				}))
			}
		}
	}
	if c.Action == nil {
		c.Help()
		return nil
	}
	return c.Action(args, c)
}
