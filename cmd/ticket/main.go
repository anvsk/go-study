package main

import (
	"go-study/pkg/util"
	"go-study/service/ticket"
)

func main() {
	util.InitUtil()
	ticket.Bootstrap()
}
