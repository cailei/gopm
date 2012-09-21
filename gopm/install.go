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
    "github.com/cailei/gopm_index/gopm/index"
    "log"
    "os/exec"
    "strings"
)

func cmd_install(args []string) int {
    known, unknown := extract_unknown_flags(args)

    var help bool
    f := flag.NewFlagSet("install_flags", flag.ExitOnError)
    f.BoolVar(&help, "help", false, "show help info")
    f.BoolVar(&help, "h", false, "show help info")
    f.Usage = print_install_help
    f.Parse(known)

    if help {
        print_install_help()
        return 0
    }

    names := f.Args()
    if len(names) == 0 {
        fmt.Println("Please specify the packages you want to install.")
        return -1
    }

    db := openLocalDB()

    var exact_matches []*index.PackageMeta
    var partial_matches []*index.PackageMeta

    for pkg, err := db.FirstPackage(); pkg != nil; pkg, err = db.NextPackage() {
        if err != nil {
            log.Fatalln(err)
        }

        low_pkg_name := strings.ToLower(pkg.Name)

        for _, name := range names {
            low_name := strings.ToLower(name)
            if low_pkg_name == low_name {
                exact_matches = append(exact_matches, pkg)
            } else if strings.Contains(low_pkg_name, low_name) {
                partial_matches = append(partial_matches, pkg)
            }
        }
    }

    if len(partial_matches) > 0 {
        fmt.Println("Some packages cannot find a match, did you mean:")
        for _, p := range partial_matches {
            fmt.Printf("\t%v\n", p.Name)
        }
        return -1
    }

    if len(exact_matches) == 0 {
        fmt.Printf("Found nothing!\n")
    }

    var failed_pkgs []string
    for _, p := range exact_matches {
        succ := install_package(p, unknown)
        if !succ {
            failed_pkgs = append(failed_pkgs, p.Name)
        }
    }

    if len(failed_pkgs) > 0 {
        fmt.Println("These packages failed to install:")
        for _, name := range failed_pkgs {
            fmt.Printf("\t%v\n", name)
        }
        return -1
    }

    return 0
}

func install_package(pkg *index.PackageMeta, args []string) bool {
    fmt.Printf("Installing %v\n", pkg.Name)

    // try installing from repos
    succ := false
    for _, repo := range pkg.Repositories {
        succ = install_from_repo(repo, args)
        if succ {
            break
        }
    }

    if !succ {
        fmt.Println("Installation failed on all repos")
    }

    return succ
}

func install_from_repo(repo string, args []string) bool {
    var cmd_line []string
    cmd_line = append(cmd_line, []string{"go", "get"}...)
    cmd_line = append(cmd_line, args...)
    cmd_line = append(cmd_line, repo)
    //fmt.Printf("Running command: %v\n", cmd_line)
    cmd := exec.Command("go", cmd_line...)
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

func extract_unknown_flags(args []string) (known []string, unknown []string) {
    for _, a := range args {
        if strings.HasPrefix(a, "-") {
            if a == "-h" || a == "--help" {
                known = append(known, a)
            } else {
                unknown = append(unknown, a)
            }
        } else {
            known = append(known, a)
        }
    }
    return
}

func print_install_help() {
    fmt.Print(`
gopm install <pkg1> [pkg2...]:
    install packages.

    this command will invoke 'go get' to get the job done, you may specify any valid 'go get' flags (see 'go get -h') in the command line, and they will be passed directly to the 'go get'.

e.g.
    'gopm install -u <package>'     install or upgrade a package
    'gopm install -a <package>'     install or upgrade a package

options:
    -h, -help       show help info

`)
}
