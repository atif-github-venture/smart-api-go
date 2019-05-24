package model

import (
	"gopkg.in/mgo.v2"
)

type ObjectRepo struct {
	Identifier  string `json:"identifier"`
	Property    string `json:"property"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdby"`
}

func EnsureIndex(dbname string, collectioname string, s *mgo.Session) *mgo.Session {
	session := s.Copy()
	c := session.DB(dbname).C(collectioname)
	for _, key := range []string{"identifier", "property"} {
		index := mgo.Index{
			Key:    []string{key},
			Unique: true,
		}
		if err := c.EnsureIndex(index); err != nil {
			panic(err)
		}
	}
	//index := mgo.Index{
	//	Key:      []string{"identifier", "property"},
	//	Unique:   true,
	//	DropDups: true,
	//}
	//err := c.EnsureIndex(index)
	//if err != nil {
	//	panic(err)
	//}
	return session
}
