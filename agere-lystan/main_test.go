package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"smartapigo/agere-lystan/app"
	"smartapigo/agere-lystan/config"
	"testing"
)

var a app.App
var basePath = "/agere-lystan"
var query = "?action=agere-lystan"
var notAvailQuery = "?action=agere-lystan1"
var payload = []byte(`{"action": "agere-lystan"}`)
var badPayloadJson = []byte(`{"action": "agere-lystan"...}`)

func TestMain(m *testing.M) {
	config := config.GetConfig()
	a.Initialize(config)
	code := m.Run()
	os.Exit(code)
}

//item not found
func TestGetObjNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", basePath, nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

//corrupt json body
func TestAddObjBadRequestJsonBody(t *testing.T) {
	req, _ := http.NewRequest("POST", basePath, bytes.NewBuffer(badPayloadJson))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

//success
func TestAddObj(t *testing.T) {
	req, _ := http.NewRequest("POST", basePath, bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
}

//general
func TestGetObj(t *testing.T) {
	req, _ := http.NewRequest("GET", basePath, nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

//duplicate post
func TestAddObjBadRequest(t *testing.T) {
	req, _ := http.NewRequest("POST", basePath, bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

//bad url
func TestDeleteObjNotFoundBadURL(t *testing.T) {
	req, _ := http.NewRequest("DELETE", basePath, bytes.NewBuffer(nil))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

//item not found
func TestDeleteObjNotFound(t *testing.T) {
	req, _ := http.NewRequest("DELETE", basePath+notAvailQuery, bytes.NewBuffer(nil))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

//success
func TestDeleteObj(t *testing.T) {
	req, _ := http.NewRequest("DELETE", basePath+query, bytes.NewBuffer(nil))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNoContent, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf(":) :) :) Oops!!!! Expected response code %d. Got %d\n", expected, actual)
	}
}
