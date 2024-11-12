package contas

import "go_alura_oo/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valor_do_saque float64) string {
	pode_sacar := valor_do_saque > 0 && valor_do_saque <= c.saldo
	if pode_sacar {
		c.saldo -= valor_do_saque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valor_do_deposito float64) (string, float64) {
	if valor_do_deposito < 0 {
		return "Depósito não realizado.", c.saldo
	}
	c.saldo += valor_do_deposito
	return "Depósito realizado.", c.saldo
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
