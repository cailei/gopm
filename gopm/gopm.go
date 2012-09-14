package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := "help"
	if len(os.Args) >= 2 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "help":
		print_help()
	case "update":
		update(os.Args)
	}
}

func print_help() {
	fmt.Print(`
Usage: gopm <command> <options>

where <command> is one of:

    update      update package index to the latest
    search      search for packages
    install     install a package
    version     show gopm version info

gopm help <command>
    show command specific options

`)
}
