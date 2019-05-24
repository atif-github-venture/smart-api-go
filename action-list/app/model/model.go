package model

import (
	"gopkg.in/mgo.v2"
)

type ObjectRepo struct {
	Action      string `json:"action"`
	DataDrivern bool   `json:"datadriven"`
	CreatedBy   string `json:"createdby"`
}

func EnsureIndex(dbname string, collectioname string, s *mgo.Session) *mgo.Session {
	session := s.Copy()
	c := session.DB(dbname).C(collectioname)
	index := mgo.Index{
		Key:      []string{"action"},
		Unique:   true,
		DropDups: true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	return session
}
