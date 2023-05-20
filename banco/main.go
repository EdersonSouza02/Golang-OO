package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	contaDenis := contas.ContaPoupanca{}

	contaDenis.Depositar(100)

	PagarBoleto(&contaDenis, 60)

	fmt.Println(contaDenis.Obtersaldo())

	contaEderson := contas.ContaCorrente{}

	contaEderson.Depositar(100)

	PagarBoleto(&contaEderson, 90)

	fmt.Println(contaEderson.Obtersaldo())
}

func PagarBoleto(contas verificarConta, valorDoBoleto float64) {

	contas.Sacar(valorDoBoleto)

}

type verificarConta interface {
	Sacar(valor float64) string
}
