package main

import (
    "flag"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
)

var remote_index_url string = "http://localhost:8080/all"
var local_index_url string = "my_index.json"

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

    // request the index content
    response, err := http.Get(remote_index_url)
    if err != nil {
        log.Fatal(err)
    }

    file, err := os.Create(local_index_url)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // write index content to local file
    bytes, err := io.Copy(file, response.Body)
    if err != nil {
        os.Remove(local_index_url)
        log.Fatal(err)
    }

    fmt.Printf("Successfully updated index! [total bytes: %v]\n", bytes)
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
