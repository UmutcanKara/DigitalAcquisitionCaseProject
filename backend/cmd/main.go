package main

import (
	"backend/router"
	"fmt"
	"log"
)

func main() {
	hosts := map[string]struct{}{"localhost:8080": {}, "127.0.0.1:8080": {}}

	r := router.NewRouter(hosts)

	fmt.Printf("Server running on port 8080\n")
	log.Fatal(r.Run(":8080"))
}
