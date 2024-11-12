package main

import (
	"fmt"
	"go_alura_oo/clientes"
	"go_alura_oo/contas"
)

// isso é como um contrato para as contas
type verficarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verficarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	clienteGustavo := clientes.Titular{Nome: "Gustavo", CPF: "123.111.123-12", Profissao: "Desenvolvedor"}
	contaDoGustavo := contas.ContaCorrente{Titular: clienteGustavo}

	contaDoGustavo.Depositar(100)

	fmt.Println(contaDoGustavo.GetSaldo())
	fmt.Println("----")

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(500)
	fmt.Println(contaDoDenis)

	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis.GetSaldo())
}
