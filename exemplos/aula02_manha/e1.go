package aula02manha

import (
	"fmt"

	"github.com/alexsandron3/bootcampMeli/aula1/utils"
)

func OpSoma(valor1, valor2 float64) float64 {
	fmt.Println(valor1, valor2, "zz")
	return valor1 + valor2
}

func OpSubtra(valor1, valor2 float64) float64 {
	return valor1 - valor2
}

func OpMultip(valor1, valor2 float64) float64 {
	return valor1 * valor2
}

func OpDivis(valor1, valor2 float64) float64 {

	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2
}
func retornoOperacao(valores []float64, operacao func(value1, value2 float64) float64) float64 {
	var resultado float64

	// operacoes := map[string]func(value1, value2 float64) float64{
	// 	utils.Soma:   OpSoma,
	// 	utils.Subtra: OpSubtra,
	// 	utils.Multip: OpMultip,
	// 	utils.Divis:  OpDivis,
	// }
	fmt.Println(valores, "valores")
	for i, valor := range valores {
		if i == 0 {
			resultado = valor
		} else {
			resultado = operacao(resultado, valor)
		}
	}

	return resultado
}

func OperacaoAritmetica(operador string, valores ...float64) float64 {

	switch operador {
	case utils.Soma:
		return retornoOperacao(valores, OpSoma)
	case utils.Subtra:
		return retornoOperacao(valores, OpSubtra)
	case utils.Multip:
		return retornoOperacao(valores, OpMultip)
	case utils.Divis:
		return retornoOperacao(valores, OpDivis)
	}

	return 0
}
