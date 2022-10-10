package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
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

func main() {

	conta2 := ContaCorrente{"Paulo", 111, 222, 20}
	minhaConta := ContaCorrente{}
	minhaConta.titular = "Rodrigo"
	minhaConta.numAgencia = 123321
	minhaConta.saldo = 10

	result, valorNaConta := conta2.Transferencia(10, &minhaConta)
	fmt.Println(result, valorNaConta)
	fmt.Println(conta2.saldo, " ----", minhaConta.saldo)

	testes := teste(1, 22, 33, 4, 256)
	fmt.Println(testes)
}

func teste(num ...int) []int {
	return num
}
