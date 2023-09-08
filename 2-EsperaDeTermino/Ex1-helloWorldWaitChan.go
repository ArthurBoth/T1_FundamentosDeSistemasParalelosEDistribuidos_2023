/*
	 Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
	   >>> Veja o Ex0 desta série
	   ABRE E FECHA CONCORRENCIA
	   Há várias formas de esperar o término de processos concorrentes.
	   EXERCICIOS:
	     1)  isto seria uma solução para sincronizar o final do programa ?
		 	R: Sim, mas não é a melhor solução.
	     2)  aumente para criar 10 prodessos concorrentes say(...).
	         como voce faz a espera de todos ?
	   OBS:  tente um comando de repeticao.
*/
package main

import (
	"fmt"
)

func say(s string, c chan struct{}) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
	c <- struct{}{}
}

func main() {
	fin := make(chan struct{})
	
	go say("um", fin)
	go say("dois", fin)
	go say("três", fin)
	go say("quatro", fin)
	go say("cinco", fin)
	go say("seis", fin)
	go say("sete", fin)
	go say("oito", fin)
	go say("nove", fin)
	go say("dez", fin)

	for i := 0; i < 10; i++ {
		<- fin
	}
}
