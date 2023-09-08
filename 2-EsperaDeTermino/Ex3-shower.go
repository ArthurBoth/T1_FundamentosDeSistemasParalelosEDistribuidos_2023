/* Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
   Exemplo da Internet
   EXERCICIOS:
     1) rode o programa abaixo e interprete.
        todos os valores escritos no canal são lidos?
	 R: Não, pois o canal tem um buffer de 5, quando o buffer do canal enche, o 'main' para de tentar escrever no canal
	 	e o programa acaba, sem ler os valores restantes do canal.
     2) como isto poderia ser resolvido ?
	 R: Com um canal de buffer 0, assim o 'main' esperaria o shower ler para então escrever no canal
*/
package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	go shower(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func shower(c chan int) {
	for {
		j := <-c
		fmt.Printf("%d\n", j)
	}
}
