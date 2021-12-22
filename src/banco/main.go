package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podesacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podesacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor de deposito invalido", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (string, float64) {

	if valorDaTransferencia > 0 && valorDaTransferencia <= c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.saldo += valorDaTransferencia
		return "Transferencia realizada com sucesso", c.saldo
	} else {
		return "Valor de transferencia invalido", c.saldo
	}
}

func main() {

	conta1 := ContaCorrente{titular: "João", numeroAgencia: 1, numeroConta: 2, saldo: 100}
	conta2 := ContaCorrente{titular: "Maria", numeroAgencia: 1, numeroConta: 3, saldo: 200}

	fmt.Println(conta1.saldo)
	fmt.Println(conta2.saldo)

	status, valor := conta1.Transferir(500, &conta2)
	fmt.Println(status, valor)

	fmt.Println(conta1.saldo)
	fmt.Println(conta2.saldo)

}
