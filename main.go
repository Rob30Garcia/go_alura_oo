package main

import (
	"fmt"
	"go_alura_oo/clientes"
	"go_alura_oo/contas"
)

func main() {
	clienteGustavo := clientes.Titular{Nome: "Gustavo", CPF: "123.111.123-12", Profissao: "Desenvolvedor"}
	contaDoGustavo := contas.ContaCorrente{Titular: clienteGustavo}

	contaDoGustavo.Depositar(100)

	fmt.Println(contaDoGustavo.GetSaldo())
}
