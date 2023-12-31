/* PUCRS - Fernando Dotti
   Modelagem de Filósofos em Go
   Garfos sao canais com um elemento.
   A existencia de um elemento no canal significa que o garfo está na mesa.
  
   Exercício: rode o programa.
      observe que quando um filósofo vai pegar um garfo que nao esta na mesa, ele fica aa espera.
      isto é modelado pelo canal garfo.   observe o ponto de parada.
      descreva a situação em que o programa para.
      veja o trace de deadlock gerado pelo runtime de go.
      identifique através dele em que linha do programa cada gorotina parou.
      como voce resolveria o problema encontrado ?`
	   R: O programa para quando todos os filósofos pegam um garfo da mesma direção,
	      assim, todos ficam esperando outro largar o seu garfo para poderem comer.
		  Para resolver o problema, basta ter um número maior de garfos do que filósofos,
		  assim, será impossível, cada filósofo ter exatamente um garfo e ter de esperar 
		  outro largar o seu para poder comer.
*/
package main

import (
	"fmt"
	"strconv"
)

const (
	PHILOSOPHERS = 3
	FORKS        = 4
)

func philosopher(id int, first_fork chan struct{}, second_fork chan struct{}) {
	for {
		fmt.Println(strconv.Itoa(id) + " senta")
		<-first_fork // pega
		fmt.Println(strconv.Itoa(id) + " pegou direita")
		<-second_fork
		fmt.Println(strconv.Itoa(id) + " come")
		first_fork <- struct{}{} // devolve
		second_fork <- struct{}{}
		fmt.Println(strconv.Itoa(id) + " levanta e pensa")
	}
}

func main() {
	var fork_channels [FORKS]chan struct{}
	for i := 0; i < FORKS; i++ {
		fork_channels[i] = make(chan struct{}, 1)
		fork_channels[i] <- struct{}{} // no inicio garfo esta livre
	}
	for i := 0; i < (PHILOSOPHERS); i++ {
		fmt.Println("Filosofo " + strconv.Itoa(i))
		go philosopher(i, fork_channels[i], fork_channels[(i+1)%FORKS])
	}
	<-make(chan struct{}) // bloqueia main para nao acabar.
}
