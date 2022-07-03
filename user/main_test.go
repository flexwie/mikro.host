package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mikro.host/common"
	"mikro.host/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

var dbPath = ":memory:"

func TestCreateHandler(t *testing.T) {
	Db = common.GetDb(&dbPath)

	v, err := json.Marshal(models.CreateRequest{Name: "Felix", Mail: "test@mf.de"})
	req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(v))
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	CreateHandler().ServeHTTP(rr, req)

	body, err := ioutil.ReadAll(rr.Body)
	assert.Nil(t, err)

	response := models.CreateResponse{}
	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	assert.Equalf(t, "", response.Err, "expected empty but got %s", response.Err)
	assert.Equalf(t, "Felix", response.Value.Name, "expected Felix but got %s", response.Value.Name)
}

func TestGetGetAll(t *testing.T) {
	Db = common.GetDb(&dbPath)
	Db.Model(&models.User{}).Create(&models.User{
		Name: "Test",
		Mail: "test",
	})

	req, err := http.NewRequest(http.MethodPost, "/get-all", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	GetAllHandler().ServeHTTP(rr, req)

	body, err := ioutil.ReadAll(rr.Body)
	assert.Nil(t, err)

	response := models.GetAllResponse{}
	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	assert.Equalf(t, "", response.Err, "expected empty but got %s", response.Err)
	assert.Len(t, response.Value, 1)
}

func TestGetGetOne(t *testing.T) {
	Db = common.GetDb(&dbPath)
	var testUser models.User
	Db.Model(&models.User{}).Create(&models.User{
		Name: "Test",
		Mail: "test",
	}).First(&testUser, "name = 'Test'")
	assert.NotEqualValues(t, "", testUser.Name)
	assert.NotNil(t, testUser)

	v, err := json.Marshal(models.GetOneRequest{Id: testUser.ID})
	req, err := http.NewRequest(http.MethodGet, "/by-id", bytes.NewBuffer(v))
	//assert.Nilf(t, err, "expected nil but got %s", err.Error())

	rr := httptest.NewRecorder()
	GetOneHandler().ServeHTTP(rr, req)

	body, err := ioutil.ReadAll(rr.Body)
	assert.Nil(t, err)

	fmt.Println(string(body))

	response := models.GetOneResponse{}
	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	assert.Equalf(t, "", response.Err, "expected empty but got %s", response.Err)
	assert.Equal(t, testUser.ID, response.Value.ID)
}
