package main

import (
	"fmt"
	"os"
)

type command_fun func(args []string)

var cmd_map map[string]command_fun

func init() {
	cmd_map = map[string]command_fun{
		"update":  cmd_update,
		"search":  cmd_search,
		"install": cmd_install,
	}
}

func main() {
	args := os.Args

	// invoked without command
	if len(args) == 1 {
		print_usage()
		return
	}

	// invoked with 'help'
	if len(args) == 2 && (args[1] == "help" || args[1] == "-h") {
		print_usage()
		return
	}

	cmd := args[1]
	fun, cmd_is_valid := cmd_map[cmd]

	// invoked with a invalid command
	if !cmd_is_valid {
		fmt.Printf("\n'%v' is not a valid command, see 'gopm help'\n\n", cmd)
		return
	}

	// call the command's function
	fun(args[1:])
}

func print_usage() {
	fmt.Print(`
Usage: 'gopm <command> <options>'

where <command> is one of:

    update      update package index to the latest
    search      search for packages
    install     install a package
    version     show gopm version info

'gopm help <command>':
    show command specific options

'gopm help' or 'gopm -h':
    show this help info

`)
}
