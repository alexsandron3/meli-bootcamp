package main

import (
	"fmt"
)

const MINUTES_IN_HOUR = 60

type Usuario struct {
	Nome      string
	Sobrenome string
	Idade     int
	Email     string
	Senha     string
	// Exercicio2
	Produtos []Produto
}

// Exercicio2
type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

// Exercício3

type Servico struct {
	Nome               string
	Preco              float64
	MinutosTrabalhados int
}

type Manutencao struct {
	Nome  string
	Preco float64
}

func (u *Usuario) MudarNome(nome, sobrenome string) {
	u.Nome = nome
	u.Sobrenome = sobrenome
}

func (u *Usuario) MudarIdade(idade int) {
	u.Idade = idade
}

func (u *Usuario) MudarEmail(email string) {
	u.Email = email
}

func (u *Usuario) MudarSenha(senha string) {
	u.Senha = senha
}

// Exercicio2
func (p *Produto) NovoProduto(nome string, preco float64) *Produto {
	return &Produto{
		Nome:  nome,
		Preco: preco,
	}
}

func AdicionarProduto(u *Usuario, p *Produto, quantidade int) {
	p.Quantidade = quantidade

	u.Produtos = append(u.Produtos, *p)
}

func DeletarProdutos(u *Usuario) {
	u.Produtos = make([]Produto, 0)
}

// Exercicio3
func SomarProdutos(produtos []Produto, valorTotal *float64, done chan bool) {
	var soma float64
	for _, produto := range produtos {
		soma += produto.Preco * float64(produto.Quantidade)
	}
	*valorTotal = soma
	done <- true
}

func SomarServicos(servico []Servico, valorTotal *float64, done chan bool) {
	var soma float64
	for _, servico := range servico {
		horasTrabalhadas := float64(servico.MinutosTrabalhados) / MINUTES_IN_HOUR
		horasTrabalhadasFormatada := float64(int(horasTrabalhadas*2)) / 2
		soma += servico.Preco * horasTrabalhadasFormatada
	}
	*valorTotal = soma
	done <- true

}
func SomarManutencao(manuntencoes []Manutencao, valorTotal *float64, done chan bool) {
	var soma float64
	for _, manutencao := range manuntencoes {
		soma += manutencao.Preco
	}
	*valorTotal = soma
	done <- true

}
func main() {
	usuario := Usuario{}

	// Exercício 2
	produto := Produto{}

	// Exercício 3
	var servicos []Servico
	var manutencoes []Manutencao
	servicos = append(servicos,
		Servico{
			Nome:               "Serviço 1",
			Preco:              100,
			MinutosTrabalhados: 60,
		},
	)

	manutencoes = append(manutencoes,
		Manutencao{
			Nome:  "Manutenção 1",
			Preco: 100,
		},
	)

	usuario.MudarNome("Alexsandro", "Xavier")
	usuario.MudarIdade(24)
	usuario.MudarEmail("alexsandro.xaviers@mercadolivre.com")
	usuario.MudarSenha("1234")

	// Exercício 2
	novoProduto := produto.NovoProduto("Computador", 100)

	AdicionarProduto(&usuario, novoProduto, 1)

	fmt.Println("Usuário com produtos", usuario)
	// DeletarProdutos(&usuario)
	// fmt.Println("Usuário após deletar produtos", usuario)

	// Exercício 3
	var totalProdutos, totalServicos, totalManutencoes, total float64
	done := make(chan bool)
	go SomarProdutos(usuario.Produtos, &totalProdutos, done)
	go SomarServicos(servicos, &totalServicos, done)
	go SomarManutencao(manutencoes, &totalManutencoes, done)

	for range done {
		if len(done) == 0 {
			break
		}
	}

	total = totalProdutos + totalServicos + totalManutencoes
	fmt.Println("Total de produtos", totalProdutos)
	fmt.Println("Total dos servicos", totalServicos)
	fmt.Println("Total de manutenções", totalManutencoes)
	fmt.Println("Total:", total)

}
