package main

import (
    "flag"
    "fmt"
)

func cmd_publish(args []string) {
    // parse flags
    var help bool
    var verbose bool

    f := flag.NewFlagSet("update_flags", flag.ExitOnError)
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

    /*

       // check folder

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
    */
}

func print_publish_help() {
    fmt.Print(`
gopm publish <package folder>:
    publish your package to the central index database.

options:
    -v, -verbose    verbose
    -h, -help       show help info

`)

}
