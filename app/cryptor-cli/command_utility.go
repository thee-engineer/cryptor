package main

import (
	"fmt"
	"os"

	"cpl.li/go/cryptor"
	"github.com/fatih/color"
)

const helpMsg = `
cryptor-cli, The Cryptor command line interface allows for the management
of nodes, peers keys and all other aspects of the Cryptor package.
`

// utility help function
func help(argv []string) error {
	// display help message
	fmt.Println(helpMsg)

	// list commands
	fmt.Printf("The commands are:\n\n")
	for name, cmd := range commands {
		color.New(color.FgGreen).Fprintf(os.Stdout, "    %-10s", name)
		color.New(color.FgYellow).Fprintf(os.Stdout, "%s\n", cmd.description)
	}

	// extended help instructions
	fmt.Printf("\nUse `%s` for more information about a command.\n",
		color.YellowString("help <command>"))

	return nil
}

// utility version function
func version(argv []string) error {

	fmt.Printf("    %-10s %s\n",
		color.GreenString("cryptor pkg"),
		color.YellowString(cryptor.Version))
	fmt.Printf("    %-10s %s\n",
		color.GreenString("cryptor cli"),
		color.YellowString(Version))

	return nil
}