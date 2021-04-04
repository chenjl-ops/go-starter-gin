package main

import (
	"go-starter-gin/internal/app/server"
	"log"

)

func main() {
	err := server.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
