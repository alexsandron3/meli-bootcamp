package main

import (
	"fmt"
	"math/rand"
	"os"
)

func GenerateFileId() *string {
	id := fmt.Sprintf("%d", rand.Intn(100))
	return &id
}

func readFile() {
	_, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado ")
	}
}

func readCustomers() {
	_, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("erro: o arquivo indicado não foi encontrado ou está danificado")
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	readFile()
	fileId := GenerateFileId()

	if fileId == nil {
		panic("problema ao gerar id")
	}

	readCustomers()

}
