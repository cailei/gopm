package main

import (
    "fmt"
    "os"
)

type command_fun func(args []string)

var cmd_map = map[string]command_fun{
    "update":  cmd_update,
    "search":  cmd_search,
    "show":    cmd_show,
    "install": cmd_install,
    "create":  cmd_create,
    "publish": cmd_publish,
}

func main() {
    args := os.Args

    // invoked without command
    if len(args) == 1 {
        print_usage()
        return
    }

    // invoked with 'help' command
    if args[1] == "help" || args[1] == "-h" {
        if len(args) == 2 {
            print_usage()
            return
        } else {
            // convert 'gopm help <cmd>' to 'gopm <cmd> -h'
            args = append(args[0:1], args[2], "-h")
        }
    }

    cmd := args[1]
    fun, cmd_is_valid := cmd_map[cmd]

    // invoked with a invalid command
    if !cmd_is_valid {
        fmt.Printf("\n'%v' is not a valid command, see 'gopm help'\n\n", cmd)
        return
    }

    // call the command's function
    fun(args[2:]) // drop the 'gopm' and <command> argument
}

func print_usage() {
    fmt.Print(`
Usage: 'gopm <command> <options>'

where <command> is one of:

    update      update package index to the latest
    search      search for packages
    install     install a package
    create      create your own package
    publish     publish your package to the central index database

'gopm help <command>':
    show command specific options

'gopm help' or 'gopm -h':
    show this help info

`)
}
