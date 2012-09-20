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

    // open local index file for read
    file, err := os.Open(local_db)
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()

    keywords := f.Args()

    // use a bufio.Reader to read the index content
    r := bufio.NewReader(file)

    for {
        // read next line
        line, err := r.ReadString('\n')

        if err != nil && err != io.EOF {
            log.Fatalln(err)
        }

        // search all keywords in this line
        all_match := true
        for _, w := range keywords {
            if !strings.Contains(line, w) {
                all_match = false
                break
            }
        }

        // all keywords is contained in this line
        if all_match {
            output_package(line)
        }

        // exit if EOF
        if err == io.EOF {
            break
        }
    }
}

func output_package(line string) {
    // construt a PackageMeta from the line
    var meta index.PackageMeta
    err := meta.FromJson([]byte(line))
    if err != nil {
        log.Fatalln(err)
    }

    // print package name and description
    fmt.Printf("%v\t%v\n", meta.Name, meta.Description)
}

func print_search_help() {
    fmt.Print(`
gopm search <keywords>:
    search for a package by name or keywords.

options:
    -h, -help       show help info

`)
}
