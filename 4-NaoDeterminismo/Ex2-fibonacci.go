/* Disciplina de Modelos de Computacao Concorrente
   Escola Politecnica - PUCRS
   Prof.  Fernando Dotti
   programa da internet - site google

   Exercício:
  		leia e entenda o funcionamento
   Atenção ao fato que o processo fibonacci fica à disposição para
   sincronizar com o main, e o main decide qual canal de fibonacci ele
   usa a cada momento.
   Se voce estuda calculo de processos, esta construção equivale
   à sincronização externa de CSP
*/
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

//   ** observe a forma de declarar um processo ativo, com "go func() { ... } ()"
//   note que valem as regras de escopo:
//   func esta definido em main, entao pode usar os canais c e quit.
//   ** observe a construção de escolha não determinística no corpo de fibonacci.
//   execute.   investigue.   pergunte!!!!!!
