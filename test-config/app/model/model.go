package model

import (
	"gopkg.in/mgo.v2"
)

type ObjectRepo struct {
	TestConfig string                 `json:"testconfig"`
	Properties map[string]interface{} `json:"properties"`
	CreatedBy  string                 `json:"createdby"`
}

func EnsureIndex(dbname string, collectioname string, s *mgo.Session) *mgo.Session {
	session := s.Copy()
	c := session.DB(dbname).C(collectioname)
	index := mgo.Index{
		Key:      []string{"testconfig"},
		Unique:   true,
		DropDups: true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	return session
}
