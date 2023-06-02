package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id              int    `json:"id"`
	TransactionCode string `json:"código da transação"`
	Coin            string `json:"moeda"`
	Value           float64 `json:"valor"`
	Issuer          string `json:"emissor"`
	Receiver        string `json:"receptor"`
	TransactionDate string `json:"data da transação"`
}

type Transactions struct {
	Transactions []Transaction `json:"Transações"`
}

func getTransactionsFromJson(transactions *Transactions) error{
	file, err := os.Open("./transactions.json")

	if err != nil {
		return err

	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(transactions)
	if err != nil {
		return err
	}
	return nil
}

func GetAll(c *gin.Context) {
	var transactions Transactions

	err := getTransactionsFromJson(&transactions)

	if err != nil {
		 c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Transactions": transactions,
	})
}

func GetByID(c *gin.Context) {
	var transactions Transactions
	requestParamId, err := strconv.Atoi(c.Param("id"))
	fmt.Println("err", err)
	err = getTransactionsFromJson(&transactions)

	if err != nil {
		 c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, transaction  := range transactions.Transactions {
		if transaction.Id == requestParamId {
			c.JSON(http.StatusOK, gin.H{
				"Transaction": transaction,
			})
			return
		}
	} 

	c.JSON(http.StatusOK, gin.H{
		"Transacion": []string{},
	})
	
}


func main() {
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Olá Alexsandro",
		})
	})

	r.GET("/transactions", GetAll)
	r.GET("/transactions/:id", GetByID)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
