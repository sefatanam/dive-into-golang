package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogBreedsJSON(t *testing.T) {
	// create a req
	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)

	// create a response
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.GetAllDogBreedsJSON)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Wrong response code; got %d, wanted 200", rr.Code)
	}
}
