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
    "bufio"
    "flag"
    "fmt"
    "github.com/cailei/gopm_index/gopm/index"
    "io"
    "log"
    "os"
    "strings"
)

func cmd_search(args []string) {
    var help bool
    f := flag.NewFlagSet("search_flags", flag.ExitOnError)
    f.BoolVar(&help, "help", false, "show help info")
    f.BoolVar(&help, "h", false, "show help info")
    f.Usage = print_search_help
    f.Parse(args)

    if help {
        print_search_help()
        return
    }

    keywords := f.Args()
    db := openLocalDB()
    pkgs := db.searchPackages(keywords)
    for _, p := range pkgs {
        // print package name and description
        fmt.Printf("%v\t%v\n", p.Name, p.Description)
    }
}

func print_search_help() {
    fmt.Print(`
gopm search <keywords>:
    search for a package by name or keywords.

options:
    -h, -help       show help info

`)
}
