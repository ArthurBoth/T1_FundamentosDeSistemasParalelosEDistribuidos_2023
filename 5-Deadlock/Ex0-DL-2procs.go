/* PUCRS - Fernando Dotti
   Modelagem de dois processos que precisam acessar concorrentemente dois
   recursos (r1 e r2).
   Exercicios:
       compreenda o programa
       execute e observe o comportamento.
       use a saida do runtime de go para identificar em que ponto cada processo para.
       explique a razão da parada.
	    R: O processo gera um Deadlock porque os processos A e B esperam um pelo outro
		   para liberar algum recurso, mas como ambos estão esperando, ambos param.
  
       como voce resolveria este problema alterando uma linha de codigo apenas ?
       (nao precisa acrescentar!!)
	    R: Fazendo com que ambos disputem o mesmo recurso na mesma ordem, assim, um
		   deles terá de esperar o outro utilizar os recursos para poder continuar,
		   evitando o Deadlock.
*/
package main

import "fmt"

func proc(s string, rx chan struct{}, ry chan struct{}) {
	for {
		<-rx
		<-ry
		rx <- struct{}{}
		ry <- struct{}{}
		fmt.Print(s)
	}
}

func main() {
	r1 := make(chan struct{}, 1)
	r2 := make(chan struct{}, 1)
	r1 <- struct{}{}
	r2 <- struct{}{}
	go proc("|", r1, r2) //  proc A
	go proc("-", r2, r1) //  proc B 
	/*		Minha sugestão de resposta
	go proc("|", r1, r2) //  proc A
	go proc("-", r1, r2) //  proc B 

			ou

	go proc("|", r2, r1) //  proc A
	go proc("-", r2, r1) //  proc B
	*/
	var blq chan struct{} = make(chan struct{})
	<-blq
}
