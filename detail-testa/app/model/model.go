package model

import (
	"gopkg.in/mgo.v2"
)

type ObjectRepo struct {
	MacroTestName string `json:"macrotestname"`
	MicroTest     []MicroTest
	CreatedBy     string `json:"createdby"`
}

type MicroTest struct {
	MicroTestName string `json:"microtestname"`
	Steps         []Steps
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
