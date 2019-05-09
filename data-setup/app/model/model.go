package model

import (
	"gopkg.in/mgo.v2"
)

type ObjectRepo struct {
	NameSpace string `json:"namespace"`
	Key       string `json:"key"`
	Random    bool   `json:"random"`
	Value     string `json:"value"`
}

func EnsureIndex(dbname string, collectioname string, s *mgo.Session) *mgo.Session {
	session := s.Copy()
	c := session.DB(dbname).C(collectioname)
	index := mgo.Index{
		Key:      []string{"namespace", "key"},
		Unique:   true,
		DropDups: true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	return session
}
