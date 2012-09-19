package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "strings"
)

func cmd_create(args []string) {
    // parse flags
    var help bool
    var force bool

    f := flag.NewFlagSet("create_flags", flag.ExitOnError)
    f.BoolVar(&help, "help", false, "show help info")
    f.BoolVar(&help, "h", false, "show help info")
    f.BoolVar(&force, "force", false, "force overwrite")
    f.BoolVar(&force, "f", false, "force overwrite")
    f.Usage = print_create_help
    f.Parse(args)

    if help {
        print_create_help()
        return
    }

    // get package folder
    jsons := f.Args()
    if len(jsons) == 0 {
        fmt.Print("\nPlease give a name for your <package>.json.\n")
        print_create_help()
        return
    }

    for i := 0; i < len(jsons); i++ {
        create_json(jsons[i], force)
    }
}

func create_json(json_name string, force bool) {
    file_name := json_name
    if !strings.HasSuffix(file_name, ".json") {
        file_name = json_name + ".json"
    }

    overwritten := false

    // check if the target file already exists
    _, err := os.Stat(file_name)
    if !os.IsNotExist(err) {
        if !force {
            log.Fatalf("Cannot create file '%v' which already exists. (use -f to overwrite)\n", file_name)
        } else {
            overwritten = true
        }
    }

    // create '<name>.json'
    json_file, err := os.Create(file_name)
    if err != nil {
        log.Fatal(err)
    }
    defer json_file.Close()

    content := `{
    "name": "",
    "description": "",
    "category": "",
    "keywords": [""],
    "author": ["", ""],
    "contributors":
    [
        ["", ""]
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

    if overwritten {
        fmt.Printf("Successfully overwritten '%v'.\n", file_name)
    } else {
        fmt.Printf("Successfully created '%v'.\n", file_name)
    }
}

func print_create_help() {
    fmt.Print(`
gopm create <package>:
    this wil create a <package.json> file containing information for your
    package, you should modify this file to fill in the fields manually, then
    run 'gopm publish <package.json>' to upload the information to the index
    server.

options:
    -f, -force      force overwrite existing file
    -h, -help       show help info

`)
}
