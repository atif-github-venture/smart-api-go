package model

import (
	"gopkg.in/mgo.v2"
)

type ObjectRepo struct {
	MacroTestName string   `json:"macrotestname"`
	MicroTests    []string `json:"microtests"`
	Tags          []string `json:"tags"`
	CreatedBy     string   `json:"createdby"`
	Status        bool
}

func EnsureIndex(dbname string, collectioname string, s *mgo.Session) *mgo.Session {
	session := s.Copy()
	c := session.DB(dbname).C(collectioname)
	index := mgo.Index{
		Key:      []string{"macrotestname"},
		Unique:   true,
		DropDups: true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	return session
}
