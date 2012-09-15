package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "github.com/cailei/gopm_index"
    "github.com/kr/pretty"
    "io/ioutil"
    "log"
    "os"
    "path"
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
    folders := f.Args()
    if len(folders) == 0 {
        fmt.Print("\nPlease provide a path to your package.\n")
        print_publish_help()
        return
    }
    folder := folders[0]

    // read package.json in the folder
    json_file_name := path.Join(folder, "package.json")
    json_file, err := os.Open(json_file_name)
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
        log.Fatalln("package.json: 'name' is empty")
    }

    pretty.Printf("%#v\n", meta)
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
