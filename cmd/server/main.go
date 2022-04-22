package main

import "github.com/willsmil/lover/internal/router"

func main() {
	router := router.NewRouter()
	router.Run(":8090")
}
