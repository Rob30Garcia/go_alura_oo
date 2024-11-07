package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valor_do_saque float64) string {
	pode_sacar := valor_do_saque > 0 && valor_do_saque <= c.Saldo
	if pode_sacar {
		c.Saldo -= valor_do_saque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valor_do_deposito float64) (string, float64) {
	if valor_do_deposito < 0 {
		return "Depósito não realizado.", c.Saldo
	}
	c.Saldo += valor_do_deposito
	return "Depósito realizado.", c.Saldo
}

func (c *ContaCorrente) Transferir(valor_da_tranferencia float64, conta_destino *ContaCorrente) bool {
	if valor_da_tranferencia > c.Saldo || valor_da_tranferencia < 0 {
		return false
	}

	c.Saldo -= valor_da_tranferencia
	conta_destino.Depositar(valor_da_tranferencia)
	return true
}
