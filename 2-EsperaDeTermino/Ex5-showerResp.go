/* Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
   >>> Veja antes Ex4 desta série.
   EXERCICIOS:
       1) esta é uma solução para a questão anterior ?
		R: Sim, pois o programa só termina quando o algo for escrito (e lido) em 'quit'
       2) o que garante que todos valores serão lidos antes do programa acabar ?
		R: O 'for' do shower, que só termina quando algo for escrito em 'quit'
*/
package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan struct{})
	go shower(ch, quit)
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	quit <- struct{}{}
}

func shower(c chan int, quit chan struct{}) {
	for {
		select {
		case j := <-c:
			fmt.Printf("%d\n", j)
		case <-quit:
			break
		}
	}
}
