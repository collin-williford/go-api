package main

import (
	"bytes"
	"encoding/json"
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

func TestGetFail(t *testing.T) {
	gin.SetMode(gin.TestMode)

}

func TestPOST(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/accounts", postAccounts)

	newAccount := Account{
		Type:          "Checking",
		AccountNumber: "123456",
		BankName:      "Citi",
		RoutingNumber: 10987,
		Balance:       2345000,
	}

	jsonValue, _ := json.Marshal(newAccount)
	req, err := http.NewRequest(http.MethodPost, "/accounts", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Counldn't create request: %v\n", err)
	}

	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)
	fmt.Println(response.Body)

	if response.Code == http.StatusCreated {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusCreated, response.Code)
	} else {
		t.Fatalf("Expected to get Status %d but instead got %d\n", http.StatusCreated, response.Code)
	}
}

func TestPUT(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.PUT("/accounts/:accountNumber", updateAccountByNumber)

	accountNum := "9988445"
	updatedAccount := Account{
		Type:          "Checking",
		AccountNumber: "9988445",
		BankName:      "Wells Fargo",
		RoutingNumber: 123456,
		Balance:       30000,
	}

	jsonValue, _ := json.Marshal(updatedAccount)
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/accounts/%s", accountNum), bytes.NewBuffer(jsonValue))
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

func TestGetByAccountNum(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/accounts/:accountNumber", getAccountByNumber)

	accountNum := "9988445"

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/accounts/%s", accountNum), nil)
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

func TestDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.DELETE("/accounts/:accountNumber", deleteAccountByNumber)

	accountNum := "9988445"

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/accounts/%s", accountNum), nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)
	fmt.Println(response.Body)

	if response.Code == http.StatusAccepted {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusAccepted, response.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusAccepted, response.Code)
	}
}
