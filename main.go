package main

import (
	"fmt"
	"go_alura_oo/contas"
)

func main() {
	// conta_da_silvia := ContaCorrente{}
	// conta_da_silvia.titular = "Silvia"
	// conta_da_silvia.saldo = 500

	// fmt.Println(conta_da_silvia.saldo)
	// fmt.Println(conta_da_silvia.sacar(400))
	// fmt.Println(conta_da_silvia.saldo)

	// status, saldo := conta_da_silvia.depositar(-500)
	// fmt.Println(status, "Saldo atual:", saldo)

	contaDaSilva := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDaSilva.Transferir(-200, &contaDoGustavo)

	fmt.Println(status)
	fmt.Println(contaDaSilva)
	fmt.Println(contaDoGustavo)
}
