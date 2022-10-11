package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	eu := clientes.Titular{Nome: "Paulo", Cpf: "123454321"}
	outraPessoa := clientes.Titular{Nome: "Paulo", Cpf: "123454321"}

	minhaConta := contas.ContaCorrente{Titular: eu, NumAgencia: 1234, NumConta: 231}
	outraConta := contas.ContaCorrente{Titular: outraPessoa, NumAgencia: 4321, NumConta: 123}

	minhaConta.Deposito(200)
	outraConta.Deposito(20)

	result, valorNaConta := minhaConta.Transferencia(10, &outraConta)
	fmt.Println(result, valorNaConta)
	fmt.Println(outraConta.GetSaldo(), " ----", minhaConta.GetSaldo())
}
