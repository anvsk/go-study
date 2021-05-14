package main

import (
    "go-ticket/pkg/util"
    "go-ticket/service/ticket"
)

func main() {
    util.InitUtil()
    ticket.Bootstrap()
}
