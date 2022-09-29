package main

import (
	"fmt"
	"golang-second-steps/contas"
)

// contaDaSilvia := ContaCorrente{}
// contaDaSilvia.titular = "Silvia"
// contaDaSilvia.saldo = 500

// fmt.Println(contaDaSilvia.saldo)
// status, valor := contaDaSilvia.Depositar(100)
// fmt.Println(status, valor)

// contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
// contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

// status := contaDaSilvia.Transferir(200, &contaDoGustavo)

// fmt.Println(status)
// fmt.Print(contaDaSilvia)
// fmt.Println(contaDoGustavo)

// contaDoJean := contas.ContaCorrente{Titular: clientes.Titular{
// 	Nome:      "Jean",
// 	CPF:       "000.000.000.00",
// 	Profissao: "DevOps"},
// 	NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}

// clienteJean := clientes.Titular{"Jean", "123.123.123.12", "DevOps"}
// contaDoJean := contas.ContaCorrente{clienteJean, 123, 123456, 100}

// contaExemplo := contas.ContaCorrente{}
// contaExemplo.Depositar(100)

// fmt.Println(contaExemplo.ObterSaldo())

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoJean := contas.ContaPoupanca{}
	contaDoJean.Depositar(100)
	PagarBoleto(&contaDoJean, 60)

	contaDaMariana := contas.ContaCorrente{}
	contaDaMariana.Depositar(500)
	PagarBoleto(&contaDaMariana, 10)

	fmt.Println(contaDoJean.ObterSaldo())
	fmt.Println(contaDaMariana.ObterSaldo())
}
