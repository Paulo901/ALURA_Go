package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func main() {

	minhaConta := ContaCorrente{"Paulo", 123456, 321, 2000}

	conta2 := new(ContaCorrente)
	conta2.titular = "Rodrigo"
	conta2.numAgencia = 123321
	conta2.saldo = 99

	fmt.Println(*conta2)
	fmt.Println("Dados da conta: ", minhaConta)

}
