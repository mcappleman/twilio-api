package main

import (
	"log"
	"os"

	"github.com/mcappleman/twilio-api/config"
	// "github.com/mcappleman/twilio-api/mongodb"
)

func main() {

	conf := config.DecodeConfig()

	file, err := os.OpenFile(conf.LOG_FILE, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		file, err = os.Create(conf.LOG_FILE)
		if err != nil {
			log.Println(conf.LOG_FILE)
			log.Println("Unable to create log file")
			log.Fatalln(err)
		}
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println("Logging started")

	os.Exit(0)

	// session := mongodb.NewSession(conf.DATABASE_URL, conf.DATABASE_NAME)

}
