package main

import (
    "bytes"
    "fmt"
    "github.com/cailei/gopm_index/gopm_index"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
)

var remote_db_host string = "http://localhost:8080"

func agent_get_full_index_reader() io.Reader {
    url := remote_db_host + "/all"
    return _get_body_reader(url)
}

func agent_package_name_exists(name string) bool {
    url := fmt.Sprintf("%v/name_exists?name=%v", remote_db_host, name)
    content := _get_body_content(url)
    exists, err := strconv.ParseBool(content)
    if err != nil {
        log.Fatalln(err)
    }
    return exists
}

func agent_upload_package(meta gopm_index.PackageMeta) {
    url := fmt.Sprintf("%v/publish", remote_db_host)

    // marshal PackageMeta to json
    json, err := meta.ToJson()
    if err != nil {
        log.Fatalln(err)
    }

    json_reader := bytes.NewReader(json)

    // create a POST request
    request, err := http.NewRequest("POST", url, json_reader)
    if err != nil {
        log.Fatalln(err)
    }
    request.ContentLength = int64(len(json))

    content := _post_data(request)
    if content == "" {
        log.Fatalln("Unknown error!")
    }

    succ, err := strconv.ParseBool(content)
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Print(succ)
}

func _post_data(request *http.Request) string {
    // POST
    response, err := http.DefaultClient.Do(request)
    if err != nil {
        log.Fatalln(err)
    }

    // check response
    if response.StatusCode != 200 {
        log.Fatalln(response.Status)
    }

    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalln(err)
    }
    return string(data)
}

func _get_body_reader(url string) io.ReadCloser {
    // GET the index content
    response, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }

    // check response
    if response.StatusCode != 200 {
        log.Fatalln(response.Status)
    }

    return response.Body
}

func _get_body_content(url string) string {
    body_reader := _get_body_reader(url)
    data, err := ioutil.ReadAll(body_reader)
    if err != nil {
        log.Fatalln(err)
    }
    return string(data)
}
