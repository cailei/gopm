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
    "fmt"
    "github.com/cailei/gopm_index/gopm/index"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
)

var remote_db_host string = "http://localhost:8080"

func agent_get_full_index_reader() io.ReadCloser {
    request := remote_db_host + "/all"
    return _get_body_reader(request)
}

func agent_upload_package(meta index.PackageMeta) {
    request := fmt.Sprintf("%v/publish", remote_db_host)

    // marshal PackageMeta to json
    json, err := meta.ToJson()
    if err != nil {
        log.Fatalln(err)
    }

    // create a POST request
    response, err := http.PostForm(request, url.Values{"pkg": {string(json)}})
    if err != nil {
        log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(response.Body)
    defer response.Body.Close()
    if err != nil {
        log.Fatalln(err)
    }

    if len(body) > 0 {
        fmt.Println(string(body))
    }

    // check response
    if response.StatusCode != 200 {
        log.Fatalln(response.Status)
    }
}

func _get_body_reader(request string) io.ReadCloser {
    // GET the index content
    response, err := http.Get(request)
    if err != nil {
        log.Fatalln(err)
    }

    // check response
    if response.StatusCode != 200 {
        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Fatalln(err)
        }

        if len(body) > 0 {
            fmt.Println(string(body))
        }

        log.Fatalln(response.Status)
    }

    return response.Body
}
