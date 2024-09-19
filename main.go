package main

import (
	"scaler/src/serve"
)

func main() {
	if err := serve.Serve(); err != nil {
		panic(err.Error())
	}
}
