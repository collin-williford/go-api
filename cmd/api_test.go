package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGET(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/accounts", getAccounts)

	req, err := http.NewRequest(http.MethodGet, "/accounts", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)
	fmt.Println(response.Body)

	if response.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, response.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, response.Code)
	}
}
