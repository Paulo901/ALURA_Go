package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float32
}

func main() {

	minhaConta := ContaCorrente{"Paulo", 123456, 321, 2000}

	fmt.Println("Dados da conta: ", minhaConta)

}
