package main

import (
    "fmt"
    "github.com/cailei/gopm_index/gopm_index"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "strconv"
)

var remote_db_host string = "http://localhost:8080"

func agent_get_full_index_reader() io.ReadCloser {
    request := remote_db_host + "/all"
    return _get_body_reader(request)
}

func agent_package_name_exists(name string) bool {
    request := fmt.Sprintf("%v/name_exists?name=%v", remote_db_host, name)
    content := _get_body_content(request)
    exists, err := strconv.ParseBool(content)
    if err != nil {
        log.Fatalln(err)
    }
    return exists
}

func agent_upload_package(meta gopm_index.PackageMeta) {
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

    // check response
    if response.StatusCode != 200 {
        log.Fatalln(response.Status)
    }

    body, err := ioutil.ReadAll(response.Body)
    defer response.Body.Close()
    if err != nil {
        log.Fatalln(err)
    }

    succ, err := strconv.ParseBool(string(body))
    if err != nil {
        log.Fatalln(err)
    }

    if !succ {
        log.Fatalln("Unknown error!")
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
        log.Fatalln(response.Status)
    }

    return response.Body
}

func _get_body_content(request string) string {
    body_reader := _get_body_reader(request)
    defer body_reader.Close()

    data, err := ioutil.ReadAll(body_reader)
    if err != nil {
        log.Fatalln(err)
    }
    return string(data)
}
