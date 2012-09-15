package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "path"
)

func cmd_create(args []string) {
    // parse flags
    var help bool
    var verbose bool

    f := flag.NewFlagSet("create_flags", flag.ExitOnError)
    f.BoolVar(&help, "help", false, "show help info")
    f.BoolVar(&help, "h", false, "show help info")
    f.BoolVar(&verbose, "verbose", false, "verbose")
    f.BoolVar(&verbose, "v", false, "verbose")
    f.Usage = print_create_help
    f.Parse(args)

    if help {
        print_create_help()
        return
    }

    // get package folder
    folders := f.Args()
    if len(folders) == 0 {
        fmt.Print("\nPlease provide target folder of your package.\n")
        print_create_help()
        return
    }
    folder := folders[0]

    // exit if the target folder already exists
    _, err := os.Stat(folder)
    if !os.IsNotExist(err) {
        log.Fatalf("Cannot create a package at '%v' cause the folder already exists.\n", folder)
    }

    // create package folder
    if err := os.MkdirAll(folder, os.ModeDir|0755); err != nil {
        log.Fatalln(err)
    }

    if verbose {
        fmt.Printf("Created folder '%v'\n", folder)
    }

    // create 'package.json'
    json_file_name := path.Join(folder, "package.json")
    json_file, err := os.Create(json_file_name)
    if err != nil {
        log.Fatal(err)
    }
    defer json_file.Close()

    content := `{
    "name": "",
    "description": "",
    "category": "",     // optional
    "keywords": [""],   // optional, for searching
    "author": {"name": "", "email": ""},
    "contributors":     // optional
    [
        {"name": "", "email": ""}
    ],
    "repositories":
    [
        ""
    ]
}
`
    _, err = json_file.Write([]byte(content))
    if err != nil {
        log.Fatalln(err)
    }

    if verbose {
        fmt.Printf("Created 'package.json'\n")
    }

    fmt.Printf("Successful!\n")

}

func print_create_help() {
    fmt.Print(`
gopm create <package>:
    create a <package.json> file containing skeleton information for your package, you should modify this file to fill in the fields manually, then
    'gopm publish <package.json>' to upload the information to the index server, make the package available to others.

options:
    -v, -verbose    verbose
    -h, -help       show help info

`)
}
