package main

import "fmt"

func main() {
	var titular string = "Gabriel"
	var numeroAgencia int = 589
	var numeroConta int = 123456
	var saldo float64 = 125.50

	var titular2 string = "Carlos"
	var numeroAgencia2 int = 589
	var numeroConta2 int = 123457
	var saldo2 float64 = 100.50

	fmt.Println(titular, numeroAgencia, numeroConta, saldo)
	fmt.Println(titular2, numeroAgencia2, numeroConta2, saldo2)
}
