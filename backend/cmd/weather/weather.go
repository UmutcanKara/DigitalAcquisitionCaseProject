package main

import (
	"backend/router"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func main() {
	pwd, _ := os.Getwd()
	environmentPath := filepath.Join(pwd, ".env")
	err := godotenv.Load(environmentPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	hosts := map[string]struct{}{"localhost:8080": {}, "127.0.0.1:8080": {}}

	// initialize new router and add handlers
	r := router.WeatherRouter(hosts)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
