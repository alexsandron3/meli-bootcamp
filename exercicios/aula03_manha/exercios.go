package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Exercicios() {
	// // Ex 1:
	// produto := []byte("Nome do produto: Água, Quantidade: 10, ID: 010101")

	// err := os.WriteFile("produtos.csv", produto, 0644)

	// if err != nil {
	// 	log.Panic(err)
	// }

	// // Ex 2:

	// file, err := os.ReadFile("./produtos.csv")

	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Printf("%s \n", file)

	produtos := [][]string{
		{"ID", "Produto", "Quantidade"},
		{"01ff0", "Água", "100"},
	}

	file, err := os.Create("produtos.csv")

	if err != nil {
		log.Fatal(err)
	}

	arquivoCsv := csv.NewWriter(file)

	for _, produto := range produtos {
		arquivoCsv.Write(produto)
	}
	file.Close()

	// produtos, err = os.Open("produtos.csv")

	if err != nil {
		log.Fatal(err)
	}
	for _, produto := range produtos {
		fmt.Println(produto)
	}
}

func main() {
	Exercicios()
}
