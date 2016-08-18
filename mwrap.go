package mwrap

import (
	"gopkg.in/mgo.v2"
	"log"
)

func NewSession(dbAddr string) (*mgo.Session, error) {
	session, err := mgo.Dial(dbAddr)
	if err != nil {
		log.Println(err)
	}

	return session, err
}

func GetColl(dbAddr string, dbName string, collName string) (*mgo.Collection, error) {
	session, err := NewSession(dbAddr)
	if err != nil {
		log.Println(err)
	}

	c := session.DB(dbName).C(collName)

	return c, err
}
