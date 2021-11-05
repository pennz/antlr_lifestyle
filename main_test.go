package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/stretchr/testify/assert"
	"github.com/pennz/antlr_lifestyle/model"
)

func TestGetThings(t *testing.T) {
	env, err := model.GetFakeEnv()
	if err != nil {
		log.Fatal(err)
	}
	pgenv, err := model.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name    string
		fileds  model.Env
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"GetDB_Fake", env, 200, false},
		{"GetDB_postgres", pgenv, 200, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupRouter()
			addModelFunc2Router(router, tt.fileds)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/things", nil)
			reqr, _ := http.NewRequest("GET", "/actions", nil)
			reqa, _ := http.NewRequest("GET", "/relations", nil)

			router.ServeHTTP(w, req)
			assert.Equal(t, tt.want, w.Code)
			router.ServeHTTP(w, reqr)
			assert.Equal(t, tt.want, w.Code)
			router.ServeHTTP(w, reqa)
			assert.Equal(t, tt.want, w.Code)
		})
	}
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
