/* por Arthur Antunes de Souza Both e Gabriel Moraes Ferreira

   Problema:
     considere um servidor que recebe pedidos por um canal (representando uma conexao)
     ao receber o pedido, sabe-se através de qual canal (conexao) responder ao cliente.
     Abaixo uma solucao sequencial para o servidor.
   Exercicio
     deseja-se tratar os clientes concorrentemente, e nao sequencialmente.
     como ficaria a solucao ?
   Veja abaixo a resposta ...
     quantos clientes podem estar sendo tratados concorrentemente ?
  
   Exercicio:
     agora suponha que o seu servidor pode estar tratando no maximo 10 clientes concorrentemente.
     como voce faria ?
	 	R: Limitamos a quantidade de requisições de cada cliente pode fazer (antes eram infinitas).
		   Modificamos a função cliente para que ela receba mais um parâmetro e avise quando terminar.
		   Forçamos o servidor a esperar um cliente terminar quando há mais de 10 clientes ativos.
		   O resto do nosso código está separado do código original. No Main, está antes de um comentário
		   e o restante está depois do Main.
*/

package main

import (
	"fmt"
	"math/rand"
)

const (
	NCL  = 100
	Pool = 10
	ClientRequisitonLimit = 10
)


type Request struct {
	v      int
	ch_ret chan int
}

// ------------------------------------
// cliente
func cliente(i int, req chan Request, done chan struct{}) {
	var v, r int
	my_ch := make(chan int)
	for i := 0; i < ClientRequisitonLimit; i++ {
		v = rand.Intn(1000)
		req <- Request{v, my_ch}
		r = <-my_ch
		fmt.Println("cli: ", i, " req: ", v, "  resp:", r)
	}
	done <- struct{}{} // libera um slot no pool
}

// ------------------------------------
// servidor
// thread de servico calcula a resposta e manda direto pelo canal de retorno informado pelo cliente
func trataReq(id int, req Request) {
	fmt.Println("                                 trataReq ", id)
	req.ch_ret <- req.v * 2
}

// servidor que dispara threads de servico
func servidorConc(in chan Request) {
	// servidor fica em loop eterno recebendo pedidos e criando um processo concorrente para tratar cada pedido
	var j int = 0
	for {
		j++
		req := <-in
		go trataReq(j, req)
	}
}

// ------------------------------------
// main


func main() {
	fim := make(chan struct{})
	treating := make(chan struct{}, Pool)
	done := make(chan struct{}, Pool)
	go applyServerLimit(treating, done, fim) // Ativa o limitador de threads
	<- done // espera o limitador de threads estar pronto

	// ------------------------------------

	fmt.Println("------ Servidores - criacao dinamica -------")
	serv_chan := make(chan Request)
	go servidorConc(serv_chan)
	for i := 0; i < NCL; i++ {
		<- treating							// Espera um slot no pool
		go cliente(i, serv_chan, done)
	}
	fim <- struct{}{}
}
//---------------------------------------------------------------------------------------------
func applyServerLimit(treating chan struct{}, done chan struct{}, fim chan struct{}){
	for i := 0; i < Pool; i++ {
		treating <- struct{}{}
	}
	done <- struct{}{} // avisa o main que o limitador de threads está pronto
	go limitPool(treating, done, fim)
}
func limitPool(treating chan struct{}, done chan struct{}, fim chan struct{}){
	for i := 0; i < NCL;i++{
		<- done				   // Espera uma requisição terminar
		treating <- struct{}{} // libera um slot no pool
	}
	end(fim)
}

func end(fim chan struct{}){ // implementado para evitar uma finalização de programa através deadlock
	<- fim
}