package aula01tarde

import "fmt"

func Exercicio1() {
	palavra := "palavra"

	fmt.Println(palavra)

	for _, letra := range palavra {
		fmt.Printf("A letra é %c\n", letra)
	}
}

func Exercicio2(age, workTime, loanValue int) {
	canMakeLoan := age >= 22 && workTime >= 1
	const LOAN_NO_TAX_MIN = 100000
	if canMakeLoan {

		fmt.Println("Pode fazer empréstimo")
		if loanValue > LOAN_NO_TAX_MIN {
			fmt.Println("Não há juros para esse valor")
		} else {
			fmt.Println("Há juros para esse valor")
		}
	} else {
		fmt.Println("Não pode fazer empréstimo")
	}
}

func Exercicio3(monthNumber int) {
	months := map[int]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}
	fmt.Println(months[monthNumber])

	// Sim, poderia ser feito com if/else e switch/case mas acho que fica mais elegante assim
}

func Exercicio4() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	employeersOlderThan21 := 0
	fmt.Println("A idade de Benjamin é", employees["Benjamin"])

	for _, age := range employees {
		if age > 21 {
			employeersOlderThan21++
		}
	}

	employees["Federico"] = 25

	delete(employees, "Pedro")

	fmt.Println("A quantidade de funcionários com mais de 21 anos é", employeersOlderThan21)
	fmt.Println("Funcionários: ", employees)
}
