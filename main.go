package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		log.Infoln("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		log.Info("Environment file was not loaded, using System Variables")
	}
	fmt.Printf("Server starting in a %s environment \n", *environment)
	UseMongo()
	InitLogic()
	InitControllers()
	defer rep.Disconnect()
}
