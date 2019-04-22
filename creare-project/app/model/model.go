package model

import (
	"gopkg.in/mgo.v2"
)

type ObjectRepo struct {
	ProjectName string `json:"projectname"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
}

func EnsureIndex(dbname string, collectioname string, s *mgo.Session) *mgo.Session {
	session := s.Copy()
	c := session.DB(dbname).C(collectioname)
	//for _, key := range []string{"projectname", "username"} {
	//	index := mgo.Index{
	//		Key:    []string{key},
	//		Unique: true,
	//	}
	//	if err := c.EnsureIndex(index); err != nil {
	//		panic(err)
	//	}
	//}
	index := mgo.Index{
		Key:      []string{"projectname", "username"},
		Unique:   true,
		DropDups: true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	return session
}
