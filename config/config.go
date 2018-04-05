package config

import (
		"os"
)

type config struct {
		LOG_FILE 		string
		DATABASE_NAME 	string
		DATABASE_URL 	string
}

func DecodeConfig() config {

		conf := config{}
	    conf.LOG_FILE = os.Getenv("LOGFILE_PATH")
	    conf.LOG_FILE = os.Getenv("MONGO_URL")
	    conf.LOG_FILE = os.Getenv("DATABASE_NAME")

		return conf

}
