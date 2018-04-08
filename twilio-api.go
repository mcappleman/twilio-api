package main

import (
	"log"
	"os"

	"github.com/mcappleman/twilio-api/app"
)

func main() {

	file, err := os.OpenFile(os.Getenv("LOG_FILE_PATH"), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		file, err = os.Create(os.Getenv("LOG_FILE_PATH"))
		if err != nil {
			log.Println(os.Getenv("LOG_FILE_PATH"))
			log.Println("Unable to create log file")
			log.Fatalln(err)
		}
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println("Logging started")

	log.Println("SOMETHING OTHER THANT LOGGING STARTED")
	a := app.App{}
	log.Println("Main before Init")
	err = a.Init()
	if err != nil {
		log.Println("Escalated to the top")
		recover()
	}
	a.Run(":8080")

}
