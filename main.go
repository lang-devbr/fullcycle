package main

import (
	"os"

	"github.com/lang-devbr/fullcycle/client"
	"github.com/lang-devbr/fullcycle/server"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		panic("invalid arg size")
	}

	if arg[0] == "server" {
		server.Start()
	}

	if arg[0] == "client" {
		client.ProcessarCotacao()
	}
}
