package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"smartapigo/identifier-reponere/app"
	"smartapigo/identifier-reponere/config"
	"testing"
)

var a app.App
var basePath = "/identifier-reponere"
var query = "?identifier=mission-impossible"
var queryUpdate = "?identifier=mission-impossible1&property=xpath%3Dproperty%23%2Flogin1"
var payload = []byte(`{"identifier":"mission-impossible","property":"xpath=property#/login"}`)
var payloadDup = []byte(`{"identifier":"mission-impossible1","property":"xpath=property#/login"}`)
var payloadDupSec = []byte(`{"identifier":"mission-impossible","property":"xpath=property#/login1"}`)
var payloadUpdate = []byte(`{"identifier":"mission-impossible1","property":"xpath=property#/login1"}`)
var badPayloadJson = []byte(`{"identifier":"mission-impossible",..."property":"some=property"}`)

func TestMain(m *testing.M) {
	config := config.GetConfig()
	a.Initialize(config)
	code := m.Run()
	os.Exit(code)
}

//item not found
func TestGetObjNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", basePath+query, nil)
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

//duplicate post
func TestAddObjBadRequestDup(t *testing.T) {
	req, _ := http.NewRequest("POST", basePath, bytes.NewBuffer(payloadDup))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

//duplicate post
func TestAddObjBadRequestDupSec(t *testing.T) {
	req, _ := http.NewRequest("POST", basePath, bytes.NewBuffer(payloadDupSec))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

//corrupt json body
func TestUpdateObjBadRequestJsonBody(t *testing.T) {
	req, _ := http.NewRequest("PUT", basePath+query, bytes.NewBuffer(badPayloadJson))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

//bad url
func TestUpdateObjNotFoundBadURL(t *testing.T) {
	req, _ := http.NewRequest("PUT", basePath, bytes.NewBuffer(payloadUpdate))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

//item not found
func TestUpdateObjNotFound(t *testing.T) {
	req, _ := http.NewRequest("PUT", basePath+queryUpdate, bytes.NewBuffer(payloadUpdate))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

//success
func TestUpdateObj(t *testing.T) {
	req, _ := http.NewRequest("PUT", basePath+query, bytes.NewBuffer(payloadUpdate))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNoContent, response.Code)
}

//bad url
func TestDeleteObjNotFoundBadURL(t *testing.T) {
	req, _ := http.NewRequest("DELETE", basePath, bytes.NewBuffer(nil))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

//item not found
func TestDeleteObjNotFound(t *testing.T) {
	req, _ := http.NewRequest("DELETE", basePath+query, bytes.NewBuffer(nil))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

//success
func TestDeleteObj(t *testing.T) {
	req, _ := http.NewRequest("DELETE", basePath+queryUpdate, bytes.NewBuffer(nil))
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
