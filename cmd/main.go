package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Account struct {
	Type          string `json:"type"`
	AccountNumber string `json:"accountnumber"`
	BankName      string `json:"bankname"`
	RoutingNumber int    `json:"routingnumber"`
	Balance       int    `json:"balance"`
}

var accounts = []Account{
	{Type: "Checking", AccountNumber: "9988445", BankName: "Wells Fargo", RoutingNumber: 72456, Balance: 24500},
	{Type: "Savings", AccountNumber: "8877442", BankName: "Wells Fargo", RoutingNumber: 87432, Balance: 32500},
}

func main() {
	router := gin.Default()
	router.GET("/accounts", getAccounts)
	router.GET("/accounts/:accountNumber", getAccountByNumber)
	router.POST("/accounts", postAccounts)
	router.DELETE("/accounts/:accountNumber", deleteAccountByNumber)

	router.Run("localhost:8080")
}

func getAccounts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, accounts)
}

func postAccounts(c *gin.Context) {
	var newAccount Account

	if err := c.BindJSON(&newAccount); err != nil {
		return
	}

	accounts = append(accounts, newAccount)
	c.IndentedJSON(http.StatusCreated, newAccount)
}

func getAccountByNumber(c *gin.Context) {
	accountNumber := c.Param("accountNumber")

	for _, a := range accounts {
		if a.AccountNumber == accountNumber {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Account number not found"})
}

func deleteAccountByNumber(c *gin.Context) {
	accountNumber := c.Param("accountNumber")

	for index, num := range accounts {
		if num.AccountNumber == accountNumber {
			accounts = append(accounts[:index], accounts[index+1])
		}
	}

	c.IndentedJSON(http.StatusAccepted, accountNumber)
}
