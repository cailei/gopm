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
    "log"
    "os"
    "path"
)

var remote_db_host string = "http://localhost:8080"
var local_db_url string

func init() {
    goroot := os.Getenv("GOROOT")
    if goroot == "" {
        log.Fatalln("Environment $GOROOT was not defined, gopm need this to put index file.")
    }

    // make sure the db dir exists
    db_dir := path.Join(goroot, "misc", "gopm")
    _, err := os.Stat(db_dir)
    if os.IsNotExist(err) {
        os.MkdirAll(db_dir, 0755)
    }
    local_db_url = path.Join(db_dir, "gopm.index")
}
