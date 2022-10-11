package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular              clientes.Titular
	NumAgencia, NumConta int
	saldo                float64
}

func (c *ContaCorrente) Saque(quantia float64) (string, float64) {

	if quantia > 0 && quantia <= c.saldo {
		c.saldo -= quantia
		return "Saque realizado! Seu saldo : \n", c.saldo
	}
	return "valor indisponivel. Seu saldo : \n", c.saldo
}
func (c *ContaCorrente) Deposito(quantia float64) (string, float64) {

	if quantia > 0 {
		c.saldo += quantia
		return "Deposito realizado! Seu saldo : \n", c.saldo
	}
	return "valor indisponivel. Seu saldo : \n", c.saldo

}

func (c *ContaCorrente) Transferencia(quantia float64, contaDestino *ContaCorrente) (string, float64) {

	if quantia <= c.saldo {
		c.Saque(quantia)
		contaDestino.Deposito(quantia)
		return "Transferencia Concluida, seu saldo: \n", c.saldo
	}
	return "algo deu errado, verifique seu saldo: ", c.saldo
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
