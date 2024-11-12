package contas

import "go_alura_oo/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor_do_saque float64) string {
	pode_sacar := valor_do_saque > 0 && valor_do_saque <= c.saldo
	if pode_sacar {
		c.saldo -= valor_do_saque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valor_do_deposito float64) (string, float64) {
	if valor_do_deposito < 0 {
		return "Depósito não realizado.", c.saldo
	}
	c.saldo += valor_do_deposito
	return "Depósito realizado.", c.saldo
}

func (c *ContaCorrente) Transferir(valor_da_tranferencia float64, conta_destino *ContaCorrente) bool {
	if valor_da_tranferencia > c.saldo || valor_da_tranferencia < 0 {
		return false
	}

	c.saldo -= valor_da_tranferencia
	conta_destino.Depositar(valor_da_tranferencia)
	return true
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
