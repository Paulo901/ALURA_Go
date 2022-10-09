package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const numMonitora = 3
const frequencia = 10

func main() {

	var numero int

	//---------------------------------------

	for {
		idade, nome := apresentaUsuario()
		fmt.Println(nome, ", escolha um número: ")
		fmt.Scan(&numero)

		if numero == 1 {
			switch idade {
			case 10:
				fmt.Println("menor de idade")
			case 18:
				fmt.Println("maior de idade")
				pegaNaNet()
				fmt.Println("\n \n ")
				pegaLogs()
			default:
				fmt.Println("idade fora da regra")
			}
			//os.Exit(0)
		} else if numero == 2 {
			fmt.Println("número incorreto")
			os.Exit(-1)
		}
	}
}

func apresentaUsuario() (int, string) {
	var nome string
	linguagem := "Go"
	var idade int

	fmt.Println("Olá mundo!")
	fmt.Println("nome :")
	fmt.Scan(&nome)
	fmt.Println("idade : ")
	fmt.Scan(&idade)

	fmt.Println("Meu nome é ", nome)
	fmt.Println("Tenho", idade, "anos e estou aprendendo", linguagem)
	return idade, nome
}

func pegaNaNet() {

	sites := leArquivo()

	for cont := 0; cont < numMonitora; cont++ {
		for i, site := range sites {
			fmt.Println("testando site ", i, ": ", site)
			testaSite(site)
		}
		fmt.Println("")
		time.Sleep(frequencia * time.Second)
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Seu site está ativo \n STATUS", resp.StatusCode)
		registraLog(site, true)
	} else {
		fmt.Println("inativo ou inexistente", resp.StatusCode)
		registraLog(site, false)
	}
}

func leArquivo() []string {
	var linhas []string

	arquivo, err := os.Open("texto.txt")
	//	arquivo, err := ioutil.ReadFile("texto.txt")
	//	fmt.Println(string(arquivo))
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		linhas = append(linhas, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return linhas
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("registro.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 03:04:05") + " - " + site + " _ONLINE: " + strconv.FormatBool(status) + "\n")

}

func pegaLogs() {

	arquivo, err := ioutil.ReadFile("registro.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	fmt.Println(string(arquivo))
}
