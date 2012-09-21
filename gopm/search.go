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
    "flag"
    "fmt"
    // "github.com/cailei/gopm_index/gopm/index"
    // "io"
    "log"
    // "os"
    "strings"
)

func cmd_search(args []string) int {
    var help bool
    f := flag.NewFlagSet("search_flags", flag.ExitOnError)
    f.BoolVar(&help, "help", false, "show help info")
    f.BoolVar(&help, "h", false, "show help info")
    f.Usage = print_search_help
    f.Parse(args)

    if help {
        print_search_help()
        return 0
    }

    keywords := f.Args()
    db := openLocalDB()

    for pkg, err := db.FirstPackage(); pkg != nil; pkg, err = db.NextPackage() {
        if err != nil {
            log.Fatalln(err)
        }

        text := pkg.Name + "." + strings.Join(pkg.Keywords, ".") + "." + pkg.Description

        all_match := true
        for _, w := range keywords {
            if !strings.Contains(text, w) {
                all_match = false
                break
            }
        }

        if all_match {
            // print package name and description
            fmt.Printf("%v\t%v\n", pkg.Name, pkg.Description)
        }
    }

    return 0
}

func print_search_help() {
    fmt.Print(`
gopm search <keywords>:
    search for a package by name or keywords.

options:
    -h, -help       show help info

`)
}
