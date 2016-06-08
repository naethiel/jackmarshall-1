package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"camlistore.org/pkg/test/dockertest"

	"github.com/julienschmidt/httprouter"

	"gopkg.in/mgo.v2"
)

func TestNewCreateScenarioHandler(t *testing.T) {

	containerID, ip := dockertest.SetupMongoContainer(t)
	defer containerID.KillRemove(t)

	database, err := mgo.Dial(ip)
	if err != nil {
		log.Fatalln(err)
	}
	defer database.Close()

	router := httprouter.New()

	router.POST("/scenario", NewCreateScenarioHandler(database))

	server := httptest.NewServer(router)
	defer server.Close()

	testCases := []struct {
		json       string
		statusCode int
	}{
		{
			json: `{
				"name": "ScenarioTest",
				"year": 2016,
				"link": "http://google.com"
			}`,
			statusCode: 200,
		},
		{
			json:       `{}`, //body empty
			statusCode: 400,
		},
		{
			json: `{
				"name": "",
				"year": 2016,
				"link": "http://google.com"
			}`, //empty name
			statusCode: 400,
		},
		{
			json: `{
				"name": "ScenarioTest",
				"year": false,
				"link": "http://google.com"
			}`, //body with invalid data
			statusCode: 400,
		},
		{
			json: `{
				"name": "ScenarioTest",
				"year": 2016,
				"link": "http://google.com",
				"moredata": "some useless data"
			}`, //body with too many parameters
			statusCode: 200,
		},
	}

	for i, c := range testCases {

		request, err := http.NewRequest("POST", server.URL+"/scenario", strings.NewReader(c.json))
		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Log("case ", i, " - Unexpected error : ", err)
			t.Fail()
		}

		if res.StatusCode != c.statusCode {
			t.Log("case ", i, " - Unexpected result")
			t.Log("expected : ", c.statusCode)
			t.Log("but was  : ", res.StatusCode)
			t.Fail()
		} else {
			t.Log("case ", i, " - Succes")
		}
	}
}
