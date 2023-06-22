package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaGabriel := ContaCorrente{titular: "Gabriel", numeroAgencia: 5441, numeroConta: 1234, saldo: 122.50}
	contaCarlos := ContaCorrente{"Carlos", 5443, 1233, 12.50}

	fmt.Println(contaGabriel)
  fmt.Println(contaCarlos)
}
