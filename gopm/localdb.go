package main

import (
    "bufio"
    "fmt"
    "github.com/cailei/gopm_index/gopm/index"
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
    db.file, err = os.Open(local_db)
    if err != nil {
        log.Fatalln(err)
    }

    // use a bufio.Reader to read the index content
    db.reader = bufio.NewReader(file)

    runtime.SetFinalizer(&db, func(db *LocalDB) { db.file.Close() })

    return db
}

func (db *LocalDB) searchPackages(keywords []string, match_pred MatchPred) []index.PackageMeta {
    return db.matchPackages(keywords, matchPredForSearch)
}

func (db *LocalDB) getPackagesByNames(names []string) []index.PackageMeta {
    pkgs := db.matchPackages(keywords, matchPredForInstall)
    if len(pkgs) == 0 {
        pkgs = db.matchPackages(keywords, matchPredForSimilarName)
        if len(pkgs) == 0 {
            fmt.Printf("Cannot find packages\n")
        } else {
            fmt.Printf("Did you mean")
        }
    }
}

type MatchPred func(meta index.PackageMeta, keywords []string) bool

func (db *LocalDB) matchPackages(keywords []string, match_pred MatchPred) []index.PackageMeta {
    db.file.Seek(0, os.SEEK_SET)

    var pkgs []index.PackageMeta

    for {
        // read next line
        line, err := db.reader.ReadString('\n')

        if err != nil && err != io.EOF {
            log.Fatalln(err)
        }

        // exit if EOF
        if err == io.EOF {
            break
        }

        // construct a PackageMeta from the line
        var meta index.PackageMeta
        err = meta.FromJson([]byte(line))
        if err != nil {
            log.Fatalln(err)
        }

        // search for all keywords in package name and description
        if match_pred(meta, keywords) {
            pkgs = append(pkgs, meta)
        }
    }

    return pkgs
}

func matchPredForSearch(meta index.PackageMeta, keywords []string) bool {
    text := meta.Name + "." + meta.Description
    all_match := true
    for _, w := range keywords {
        if !strings.Contains(text, w) {
            all_match = false
            break
        }
    }
    return all_match
}

func matchPredForInstall(meta index.PackageMeta, keywords []string) bool {
    text := meta.Name
    all_match := true
    for _, w := range keywords {
        if text != w {
            all_match = false
            break
        }
    }
    return all_match
}

func matchPredForSimilarName(meta index.PackageMeta, keywords []string) bool {
    text := meta.Name
    all_match := true
    for _, w := range keywords {
        if !strings.Contains(text, w) {
            all_match = false
            break
        }
    }
    return all_match
}
