package model

import (
	"gopkg.in/mgo.v2"
)

type ObjectRepo struct {
	MicroTestName string `json:"microtestname,omitempty"`
	Steps        [] Steps
	CreatedBy    string `json:"createdby,omitempty"`
}

type Steps struct {
	Identifier string `json:"identifier,omitempty"`
	Action     string `json:"action,omitempty"`
	Data       string `json:"data,omitempty"`
}

func EnsureIndex(dbname string, collectioname string, s *mgo.Session) *mgo.Session {
	session := s.Copy()
	c := session.DB(dbname).C(collectioname)
	index := mgo.Index{
		Key:      []string{"microtestname"},
		Unique:   true,
		DropDups: true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	return session
}
