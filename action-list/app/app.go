package app

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"smartapigo/action-list/app/handler"
	"smartapigo/action-list/config"
)

type App struct {
	Router  *mux.Router
	MongoDB *mgo.Session
	conf    *config.Config
}

func (a *App) Initialize(config *config.Config) {
	dbURI := config.DB.MongoUrl
	a.conf = config
	session, err := mgo.Dial(dbURI)
	if err != nil {
		panic(err)
	}
	a.MongoDB = session
	a.MongoDB.SetMode(mgo.Monotonic, true)
	a.Router = mux.NewRouter().StrictSlash(true)
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/action-list", a.GetObject)
	a.Post("/action-list", a.AddObject)
	a.Delete("/action-list", a.DeleteObject)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) GetObject(w http.ResponseWriter, r *http.Request) {
	handler.GetObjects(a.conf, a.MongoDB, w, r)
}

func (a *App) AddObject(w http.ResponseWriter, r *http.Request) {
	handler.AddObj(a.conf, a.MongoDB, w, r)
}

func (a *App) DeleteObject(w http.ResponseWriter, r *http.Request) {
	handler.DeleteObj(a.conf, a.MongoDB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
