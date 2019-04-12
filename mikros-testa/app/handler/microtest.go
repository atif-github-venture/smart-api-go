package handler

import (
	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"smartapigo/mikros-testa/app/model"
	"smartapigo/mikros-testa/config"
)

func GetObjects(co *config.Config, s *mgo.Session, w http.ResponseWriter, r *http.Request) {
	session := s.Copy()
	m := make(map[string]string)
	var objrep []model.ObjectRepo
	err := errors.New("yoyo")
	c := session.DB(co.DB.Name).C(co.DB.Collection)
	values := r.URL.Query()

	if len(values) > 0 {
		for k, v := range values {
			m[k] = v[0]
		}
		err = c.Find(m).All(&objrep)
	} else {
		err = c.Find(bson.M{}).All(&objrep)
	}

	if err != nil {
		ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed to get any object: ", err)
		return
	}
	if objrep == nil {
		ErrorWithJSON(w, "Not found", http.StatusNotFound)
		//log.Println("Object queried not available: ", err)
		return
	}

	respBody, err1 := json.MarshalIndent(objrep, "", "  ")
	if err1 != nil {
		log.Fatal(err1)
	}

	ResponseWithJSON(w, respBody, http.StatusOK)
}

func AddObj(co *config.Config, s *mgo.Session, w http.ResponseWriter, r *http.Request) {
	session := s.Copy()
	var objrep model.ObjectRepo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&objrep)
	if err != nil {
		ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	c := session.DB(co.DB.Name).C(co.DB.Collection)

	err = c.Insert(objrep)
	if err != nil {
		if mgo.IsDup(err) {
			ErrorWithJSON(w, "This micro test already exists", http.StatusBadRequest)
			return
		}

		ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		//log.Println("Failed to insert the object: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	ResponseWithJSON(w, []byte(nil), http.StatusCreated)

}

func UpdateObj(co *config.Config, s *mgo.Session, w http.ResponseWriter, r *http.Request) {
	session := s.Copy()
	var objrep model.ObjectRepo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&objrep)
	if err != nil {
		ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	m := make(map[string]string)
	c := session.DB(co.DB.Name).C(co.DB.Collection)
	values := r.URL.Query()
	if len(values) > 0 {
		for k, v := range values {
			m[k] = v[0]
		}
		err = c.Update(m, &objrep)
	} else {
		ErrorWithJSON(w, "Wrong URL query string formation", http.StatusNotFound)
	}

	if err != nil {
		switch err {
		default:
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			//log.Println("Failed update object: ", err)
			return
		case mgo.ErrNotFound:
			ErrorWithJSON(w, "Object not found", http.StatusNotFound)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
	ResponseWithJSON(w, []byte(nil), http.StatusNoContent)

}

func DeleteObj(co *config.Config, s *mgo.Session, w http.ResponseWriter, r *http.Request) {
	session := s.Copy()
	m := make(map[string]string)
	var err error
	c := session.DB(co.DB.Name).C(co.DB.Collection)
	values := r.URL.Query()

	if len(values) > 0 {
		for k, v := range values {
			m[k] = v[0]
		}
		err = c.Remove(m)
	} else {
			ErrorWithJSON(w, "Wrong URL query string formation", http.StatusBadRequest)
	}

	if err != nil {
		switch err {
		default:
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			//log.Println("Failed delete object: ", err)
			return
		case mgo.ErrNotFound:
			ErrorWithJSON(w, "Object not found", http.StatusNotFound)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
	ResponseWithJSON(w, []byte(nil), http.StatusNoContent)
}
