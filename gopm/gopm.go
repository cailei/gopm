/*
gopm (Go Package Manager)
Copyright (c) 2012 cailei (dancercl@gmail.com)

The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
    "fmt"
    "log"
    "os"
)

type command_fun func(args []string) int

var cmd_map = map[string]command_fun{
    "update":  cmd_update,
    "search":  cmd_search,
    "show":    cmd_show,
    "install": cmd_install,
    "create":  cmd_create,
    "publish": cmd_publish,
}

func main() {
    log.SetFlags(log.Lshortfile)

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
    fun, is_cmd_valid := cmd_map[cmd]

    // invoked with a invalid command
    if !is_cmd_valid {
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

    update      update your local package index
    search      search for packages
    show        show detailed information for packages
    install     install a package
    create      create your own package
    publish     publish your package to the central index database

'gopm help <command>':
    show command specific options

'gopm help' or 'gopm -h':
    show this help info

`)
}
