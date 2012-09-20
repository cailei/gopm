package main

import (
    "os"
    "runtime"
)

type LocalDB struct {
    f *os.File
}

func localdb_finalier(db *LocalDB) {
    db.f.Close()
}

func openLocalDB() *LocalDB {
    // open local index file for read
    file, err := os.Open(local_db)
    if err != nil {
        log.Fatalln(err)
    }
    runtime.SetFinalizer(&file, localdb_finalier)
}
