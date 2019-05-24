package model

import (
	"gopkg.in/mgo.v2"
)

type ObjectRepo struct {
	TestName  string   `json:"testname"`
	Tags      []string `json:"tags"`
	Steps     []Steps
	CreatedBy string `json:"createdby"`
	Status    bool
}

type Steps struct {
	Identifier string `json:"identifier"`
	Action     string `json:"action"`
	Data       string `json:"data"`
}

func EnsureIndex(dbname string, collectioname string, s *mgo.Session) *mgo.Session {
	session := s.Copy()
	c := session.DB(dbname).C(collectioname)
	index := mgo.Index{
		Key:      []string{"testname"},
		Unique:   true,
		DropDups: true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	return session
}
