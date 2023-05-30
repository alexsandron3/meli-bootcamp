package aula02tarde

import (
	"fmt"
	"time"
)

type Estudante struct {
	Nome, Sobrenome, rg string
	DataAdmissao        time.Time
}

func (e Estudante) Detalhes() {
	fmt.Println("Nome: ", e.Nome, " ", e.Sobrenome, " RG: ", e.rg, e.DataAdmissao)
}

func Exercicio1(nome, sobrenome, rg, dataAdmissao string) {

	estudante := &Estudante{
		Nome:         nome,
		rg:           rg,
		Sobrenome:    sobrenome,
		DataAdmissao: time.Now(),
	}

	estudante.Detalhes()
}

func Exercicio2() {

}

type produto struct {
	Nome  string
	Preco float64
}
type Loja struct {
	Prdutos []produto
}
type Produto interface {
	CacularCusto(porte string) float64
}
type Ecommerce interface {
	Total() float64
	Adicionar(produto)
}

func (p *produto) CacularCusto(porte string) float64 {

	return 0.0
}

func (p *produto) NovoProduto(tipo, nome string, preco float64) produto {
	produto := produto{
		Nome:  nome,
		Preco: preco,
	}
	return produto
}
