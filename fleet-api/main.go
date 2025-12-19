package main

import (
	"fleet-api/routing"
)

func main() {
	router := routing.NewRouter()
	router.Listen(":8080")
}
