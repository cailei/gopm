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
    "io"
    "io/ioutil"
    "log"
    "os"
)

var local_db string = "my_index.json"

func cmd_update(args []string) {
    // parse flags
    var help bool
    f := flag.NewFlagSet("update_flags", flag.ExitOnError)
    f.BoolVar(&help, "help", false, "show help info")
    f.BoolVar(&help, "h", false, "show help info (shorthand)")
    f.Usage = print_update_help
    f.Parse(args)

    if help {
        print_update_help()
        return
    }

    // open a temporary file to receive the index
    temp_file, err := ioutil.TempFile("", "gopm_local_db_")
    if err != nil {
        log.Fatalln(err)
    }
    defer func() {
        temp_file.Close()
        os.Remove(temp_file.Name())
    }()

    // open remote db
    agent := newAgent()
    remote_db := agent.getFullIndexReader()

    // write index content to the temp file
    _, err = io.Copy(temp_file, remote_db)
    if err != nil {
        log.Fatalln(err)
    }

    // copy temp file content to the local db file
    db_file, err := os.Create(local_db)
    if err != nil {
        log.Fatalln(err)
    }
    defer db_file.Close()

    temp_file.Seek(0, os.SEEK_SET)
    copyed_bytes, err := io.Copy(db_file, temp_file)
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Printf("Successfully updated index! [total bytes: %v]\n", copyed_bytes)
}

func print_update_help() {
    fmt.Print(`
gopm update:
    update gopm local index.

options:
    -v, -verbose    verbose
    -h, -help       show help info

`)

}
