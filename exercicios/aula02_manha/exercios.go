package aula02manha

import (
	"errors"
	"fmt"

	"github.com/alexsandron3/bootcampMeli/aula1/utils"
)

func Exercicio1(salario float64) (imposto float64) {
	imposto = salario * 0.17

	if salario > 150000 {
		imposto = salario * 0.27
	}
	return
}

func Exercicio2(notas ...int) (int, error) {
	var somaNotas int
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("o número não pode ser negativo")
		}
		somaNotas = nota
	}
	media := somaNotas / len(notas)
	return media, nil
}

func Exercicio3(categoria string, horasTrabalhadas int) float64 {

	var salario float64
	if categoria == "C" {
		salario = float64(horasTrabalhadas) * 1000
	}

	if categoria == "B" {
		salario = float64(horasTrabalhadas) * 1500
		if horasTrabalhadas > 160 {
			salario *= 1.20
		}
	}

	if categoria == "A" {
		salario = float64(horasTrabalhadas) * 1500
		if horasTrabalhadas > 160 {
			salario *= 1.50
		}
	}

	return salario
}

func Minimo(valores ...int) int {
	valorMinimo := valores[0]
	for _, valor := range valores {
		if valor < valorMinimo {
			valorMinimo = valor
		}
	}

	return valorMinimo
}

func Maximo(valores ...int) int {
	valorMaximo := valores[0]
	for _, valor := range valores {
		if valor > valorMaximo {
			valorMaximo = valor
		}
	}

	return valorMaximo
}

func Medio(valores ...int) int {
	fmt.Println(valores)

	var valorMedio int
	var valorTotal int

	for _, valor := range valores {
		valorTotal += valor
	}
	valorMedio = valorTotal / len(valores)

	return valorMedio

}

// Adriana Meier e Vitoria Raissa
func Exercicio4(tipo string) (func(valores ...int) int, error) {
	switch tipo {
	case utils.Minimum:
		return Minimo, nil
	case utils.Average:
		return Medio, nil
	case utils.Maximum:
		return Maximo, nil
	default:
		return nil, errors.New(" tipo de calculo não definido")
	}

}
