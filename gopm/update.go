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
    remote_db := agent_get_full_index_reader()
    defer remote_db.Close()

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
