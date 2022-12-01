package main

import (
	"github.com/bagasfathoni/go-clean-architecture-template/delivery"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./conf.env")
	if err != nil {
		panic(err)
	}

	delivery.Server().Run()
}
