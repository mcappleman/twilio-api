package main

import (
	"fmt"
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

	a := app.App{}
	err = a.Init()
	if err != nil {
		fmt.Println("Escalated to the top")
		log.Println("Escalated to the top")
		fmt.Println(err)
		log.Fatalln(err)
	}
	a.InitRoutes()
	a.Run(":8080")
	// fmt.Println("Running on 8080")

}
