package main

import (
	"fallback/internal/server"
)

func main() {
	if err := server.SetupAndListen(); err != nil {
		panic(err.Error())
	}
}
