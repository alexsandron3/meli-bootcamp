package exercicios

import "fmt"

func Exercicio1() {
	nome := "Alex"
	idade := 24

	fmt.Println("Meu nome é", nome, "e tenho", idade, "anos")
}

func Exercicio2() {
	var (
		temperatura        float32
		umidade            string
		pressaoAtmosferica uint16
	)
	temperatura = 35.5
	umidade = "alta"
	pressaoAtmosferica = 1019

	fmt.Println("A temperatura é", temperatura, "graus", "com umidade", umidade, "e pressão atmosférica de", pressaoAtmosferica, "hPa")

}

func Exercicio3() {
	var nome string
	var sobrenome string
	var idade int
	licencaParaDirigir := true
	var estaturaDaPessoa int
	quantidadeDeFilhos := 2

	fmt.Println("Meu nome é", nome, sobrenome, "e tenho", idade, "anos", "e tenho", estaturaDaPessoa, "m de altura", "e tenho", quantidadeDeFilhos, "filhos", "e tenho licença para dirigir?", licencaParaDirigir)
}
