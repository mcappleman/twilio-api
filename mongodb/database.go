package mongodb

import (
	"log"
	"gopkg.in/mgo.v2"
)

type DatabaseSession struct {
    *mgo.Session
    databaseName string
}

func NewSession(url string, name string) *DatabaseSession, error {

	url = url + "/" + name
    session, err := mgo.Dial(url)
    if err != nil {
    	log.Println("Error Connecting to Mongo Database")
        return err
    }

    return &DatabaseSession{session, name}

}

func (s *DatabaseSession) Database() *mgo.Database {

    return s.DB(s.databaseName)

}
