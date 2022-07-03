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

func TestCreateUser(t *testing.T) {
	Db = common.GetDb(&dbPath)
	defer func() {
		sqldb, _ := Db.DB()
		sqldb.Close()
	}()

	server := GetCreate()
	ts := httptest.NewServer(server)
	defer ts.Close()

	v, err := json.Marshal(models.CreateRequest{Name: "Felix", Mail: "test@mf.de"})
	res, err := http.Post(ts.URL, "", bytes.NewBuffer(v))
	assert.Nil(t, err)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
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

	server := GetGetAll()
	ts := httptest.NewServer(server)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.Nil(t, err)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
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
	}).First(&testUser, "name = Test")
	assert.NotEqualValues(t, "", testUser.Name)

	server := GetGetOne()
	ts := httptest.NewServer(server)
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/?id=%d", ts.URL, testUser.ID))
	assert.Nilf(t, err, "expected nil but got %s", err.Error())

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	assert.Nil(t, err)

	response := models.GetOneResponse{}
	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	assert.Equalf(t, "", response.Err, "expected empty but got %s", response.Err)
	assert.Equal(t, testUser.ID, response.Value.ID)
}
