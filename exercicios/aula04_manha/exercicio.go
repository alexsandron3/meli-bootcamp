package main

import (
	"fmt"
)

type salarioBaixoError struct {
	salario int
}

func (e salarioBaixoError) Error() string {
	return fmt.Sprintf("erro: O salário digitado (%d) não está dentro do valor mínimo", e.salario)
}

func ValidaSalario(salario int) error {
	if salario < 15_000 {
		return salarioBaixoError{salario}
	}
	fmt.Println("Necessário pagamento de imposto")
	return nil
}

func main() {
	salario := 10_000
	err := ValidaSalario(salario)
	if err != nil {
		fmt.Println(err)
	}
}
