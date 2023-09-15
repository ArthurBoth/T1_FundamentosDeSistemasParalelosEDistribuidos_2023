/* PUCRS - Fernando Dotti
   Exercício:
     1) que condição de coffman é quebrada com a solução abaixo ?
	  R: A de espera circular.
     2) argumente em portugues porque esta solução não tem deadlock.
	  R: A solução não possui Deadlock pois o filósofo canhoto começa pelo garfo da sua esquerda
	  	 e não pegará o da direita até ter o da esquerda, evitando assim o deadlock.
     3) voce poderia ter mais filósofos canhotos ?
	  R: Sim, desde que o sempre hajam filósofos destros e canhotos
             implemente e teste.
*/
package main

import (
	"fmt"
	"strconv"
)

const (
	LEFTHANDED = 2
	PHILOSOPHERS = 5
	FORKS        = 5
)

func philosopher(id int, first_fork chan struct{}, second_fork chan struct{}) {
	for {
		fmt.Println(strconv.Itoa(id) + " senta\n")
		<-first_fork // pega
		<-second_fork
		fmt.Println(strconv.Itoa(id) + " come\n")
		first_fork <- struct{}{} // devolve
		second_fork <- struct{}{}
		fmt.Println(strconv.Itoa(id) + " levanta e pensa \n")
	}
}

func main() {
	var fork_channels [FORKS]chan struct{}

	for i := 0; i < FORKS; i++ {
		fork_channels[i] = make(chan struct{}, 1)
		fork_channels[i] <- struct{}{} // no inicio garfo esta livre
	}
	for i := 0; i < (PHILOSOPHERS - LEFTHANDED); i++ {
		fmt.Println("Filosofo " + strconv.Itoa(i) + " destro!")
		go philosopher(i, fork_channels[i], fork_channels[(i+1)])
	}
	for i := (PHILOSOPHERS - LEFTHANDED); i < (PHILOSOPHERS); i++ {
		fmt.Println("Filosofo " + strconv.Itoa(i) + " canhoto!")
		go philosopher(i, fork_channels[i], fork_channels[(i+1)%FORKS])
	}

	<-make(chan struct{})
}
