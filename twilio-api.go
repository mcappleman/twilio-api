package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile(os.Getenv("LOG_FILE_PATH"), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		file, err = os.Create(os.Getenv("LOG_FILE_PATH"))
		if err != nil {
			log.Println(conf.LOG_FILE)
			log.Println("Unable to create log file")
			log.Fatalln(err)
		}
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println("Logging started")

	app := App{}
	err := app.Init()
	if err != nil {
		log.Println("Escalated to the top")
		recover(err)
	}
	app.Run(":8080")

}
