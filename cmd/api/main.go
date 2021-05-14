package main

import (
    "go-ticket/pkg/store/cache"
    "go-ticket/pkg/store/db"
    "go-ticket/pkg/util"
)

func main() {
    util.InitUtil()
    cache.InitCache()
    db.InitDB()

}
