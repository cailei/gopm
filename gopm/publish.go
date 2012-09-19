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
    "encoding/json"
    "flag"
    "fmt"
    "github.com/cailei/gopm_index/gopm_index"
    "io/ioutil"
    "log"
    "os"
)

func cmd_publish(args []string) {
    // parse flags
    var help bool
    var verbose bool

    f := flag.NewFlagSet("publish_flags", flag.ExitOnError)
    f.BoolVar(&help, "help", false, "show help info")
    f.BoolVar(&help, "h", false, "show help info")
    f.BoolVar(&verbose, "verbose", false, "verbose")
    f.BoolVar(&verbose, "v", false, "verbose")
    f.Usage = print_publish_help
    f.Parse(args)

    if help {
        print_publish_help()
        return
    }

    // get package folder
    json_names := f.Args()
    if len(json_names) == 0 {
        fmt.Print("\nPlease provide a <package>.json file to publish.\n")
        print_publish_help()
        return
    }

    for i := 0; i < len(json_names); i++ {
        publish_package(json_names[i], verbose)
    }
}

func publish_package(json_name string, verbose bool) {
    // read json content
    json_file, err := os.Open(json_name)
    if err != nil {
        log.Fatalln(err)
    }
    defer json_file.Close()

    json_content, err := ioutil.ReadAll(json_file)
    if err != nil {
        log.Fatalln(err)
    }

    // unmarshal to PackageMeta object
    var meta gopm_index.PackageMeta
    if err := json.Unmarshal(json_content, &meta); err != nil {
        log.Fatalln(err)
    }

    // check mandatory fields
    if meta.Name == "" {
        log.Fatalf("%v: 'name' is empty\n", json_name)
    }
    if meta.Description == "" {
        log.Fatalf("%v: 'description' is empty\n", json_name)
    }
    if meta.Author[0] == "" {
        log.Fatalf("%v: 'author.name' is empty\n", json_name)
    }
    if meta.Repositories == nil {
        log.Fatalf("%v: 'repositories' is missing\n", json_name)
    }
    if len(meta.Repositories) == 0 {
        log.Fatalf("%v: 'repositories' is empty\n", json_name)
    }

    // post the json to the server
    agent_upload_package(meta)

    fmt.Printf("Successfully publish package '%v'\n", meta.Name)
}

func print_publish_help() {
    fmt.Print(`
gopm publish <package name>:
    publish your package to the central index database.

options:
    -v, -verbose    verbose
    -h, -help       show help info

`)

}
