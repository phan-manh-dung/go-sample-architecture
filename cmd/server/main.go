package main

import "go-sample-architecture/internal/routers"

func main() {
	r := routers.NewRouter()
	r.Run(":8080")
}
