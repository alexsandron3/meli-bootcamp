package main

import (
	"fmt"
	"math/rand"
	"os"
)

type File struct {
	FileName string
	Id       string
}

func GenerateFileId() *string {
	id := fmt.Sprintf("%d", rand.Intn(100))
	return &id
}

type Customer struct {
	Name        string
	Surname     string
	Rg          string
	PhoneNumber string
	Address     string
}

func readFile() {
	_, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado ")
	}
}

func readCustomers() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
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
