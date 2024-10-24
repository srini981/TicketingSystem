package main

import (
	"github.com/joho/godotenv"
	"intelXlabs/notificationService/pkg"
	"log"
)

func main() {
	log.Println("starting notificationService")
	log.Println("loading environment variables")
	godotenv.Load()
	log.Println("starting consumer")
	pkg.ConsumerInit()
}
