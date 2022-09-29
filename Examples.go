// package main

// import "fmt"

// type ContaCorrente struct {
// 	titular       string
// 	numeroAgencia int
// 	numeroConta   int
// 	saldo         float64
// }

// func main() {
// 	// contaDoJcqueiroz := ContaCorrente{titular: "Jcqueiroz",
// 	// 	numeroAgencia: 000, numeroConta: 123456789, saldo: 125.5}

// 	// contaDaMbqueiroz := ContaCorrente{"Mbqueiroz",
// 	// 	111, 1234567, 140.5}

// 	// contaDaMbqueiroz2 := ContaCorrente{"Mbqueiroz",
// 	// 	222, 1234567, 140.5}

// 	// fmt.Println(contaDoJcqueiroz)
// 	// fmt.Println(contaDaMbqueiroz == contaDaMbqueiroz2)

// 	var contaDaCris *ContaCorrente
// 	contaDaCris = new(ContaCorrente)
// 	contaDaCris.titular = "Cris"
// 	contaDaCris.saldo = 500

// 	var contaDaCris2 *ContaCorrente
// 	contaDaCris2 = new(ContaCorrente)
// 	contaDaCris2.titular = "Cris"
// 	contaDaCris2.saldo = 500

// 	// fmt.Println(&contaDaCris)
// 	// fmt.Println(&contaDaCris2)

// 	fmt.Println(*contaDaCris == *contaDaCris2)

// 	// contaDoJcqueiroz2 := ContaCorrente{titular: "Jcqueiroz",
// 	// 	numeroAgencia: 580, numeroConta: 123456789, saldo: 125.5}

// 	// fmt.Println(contaDoJcqueiroz == contaDoJcqueiroz2)
// }

// package main

// import (
// 	"fmt"
// )

// type Conta struct {
// 	saldo float64
// }

// func (c *Conta) depositarDezReais() float64 {
// 	return c.saldo + 10
// }

// func main() {
// 	contaTeste := Conta{saldo: 10}

// 	contaTeste.depositarDezReais()
// 	contaTeste.depositarDezReais()

// 	fmt.Println(contaTeste)
// }

// package main

// import (
// 	"fmt"
// )

// func SemParametro() string {
// 	return "Exemplo de função sem parâmetro!"
// }

// func UmParametro(texto string) string {
// 	return texto
// }

// func DoisParametros(texto string, numero int) (string, int) {
// 	return texto, numero
// }

// func main() {
// 	fmt.Println(SemParametro())
// 	fmt.Println(UmParametro("Exemplo de função com um parâmetro"))
// 	fmt.Println(DoisParametros("Passando dois parâmetros: essa string e o número", 10))
// }

// package main

// import (
// 	"fmt"
// )

// func Somando(numeros ...int) int {
// 	resultadoDaSoma := 0
// 	for _, numero := range numeros {
// 		resultadoDaSoma += numero
// 	}
// 	return resultadoDaSoma
// }

// func main() {
// 	fmt.Println(Somando(1))
// 	fmt.Println(Somando(1, 1))
// 	fmt.Println(Somando(1, 1, 1))
// 	fmt.Println(Somando(1, 1, 2, 4))
// }

// package main

// import (
// 	"fmt"
// )

// func mostraNomes(nomes ...string) {
// 	for _, nome := range nomes {
// 		fmt.Println(nome)
// 	}
// }

// func main() {
// 	mostraNomes("João", "Maria", "José")
// }

// package main

// import (
// 	"fmt"
// )

// type Pessoa struct {
// 	nome, sobrenome string
// }

// func (p *Pessoa) ExibirNomeCompleto() string {
// 	nomeCompleto := p.nome + " " + p.sobrenome
// 	return nomeCompleto
// }

// func main() {
// 	p1 := Pessoa{"Jean", "Queiroz"}
// 	fmt.Println(p1.ExibirNomeCompleto())
// }

// package main

// import (
// 	"fmt"
// )

// type Pessoa struct {
// 	nome, sobrenome string
// }

// func (p Pessoa) ExibirNomeCompleto() string {
// 	p.sobrenome = "Queiroz"
// 	nomeCompleto := p.nome + " " + p.sobrenome
// 	return nomeCompleto
// }

// func main() {
// 	p1 := Pessoa{"Jean", "Queiroz"}

// 	fmt.Println(p1.ExibirNomeCompleto())
// 	fmt.Println(p1.nome, p1.sobrenome)
// }