package main

import (
    "bufio"
    "github.com/cailei/gopm_index/gopm/index"
    "io"
    "log"
    "os"
    "runtime"
)

type LocalDB struct {
    file   *os.File
    reader *bufio.Reader
}

func openLocalDB() *LocalDB {
    db := new(LocalDB)

    // open local index file for read
    var err error
    db.file, err = os.Open(local_db_url)
    if err != nil {
        log.Fatalln(err)
    }

    // use a bufio.Reader to read the index content
    db.reader = bufio.NewReader(db.file)

    // auto Close() when GC
    runtime.SetFinalizer(db, func(db *LocalDB) { db.file.Close() })

    return db
}

func (db *LocalDB) FirstPackage() (pkg *index.PackageMeta, err error) {
    db.file.Seek(0, os.SEEK_SET)
    pkg, err = db.readPackage()
    return
}

func (db *LocalDB) NextPackage() (pkg *index.PackageMeta, err error) {
    pkg, err = db.readPackage()
    return
}

func (db *LocalDB) readPackage() (pkg *index.PackageMeta, err error) {
    pkg = nil

    line, err := db.readLine()
    if err != nil {
        return
    }

    // construct a PackageMeta from the line
    var meta index.PackageMeta
    err = meta.FromJson([]byte(line))
    if err != nil {
        return
    }

    pkg = &meta
    err = nil
    return
}

func (db *LocalDB) readLine() (line string, err error) {
    line, err = db.reader.ReadString('\n')

    // EOF is not an error
    if err == io.EOF {
        err = nil
    }

    return
}
