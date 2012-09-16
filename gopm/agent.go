package main

import (
    "io"
    "log"
    "net/http"
)

var remote_db_host string = "http://localhost:8080"

func agent_get_full_index_reader() io.Reader {
    reqeust := remote_db_host + "/all"

    // request the index content
    response, err := http.Get(reqeust)
    if err != nil {
        log.Fatalln(err)
    }

    // check response
    if response.StatusCode != 200 {
        log.Fatalln(response.Status)
    }

    return response.Body
}
