/* por Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
   ABRE E FECHA CONCORRENCIA
   Há várias formas de esperar o término de processos concorrentes.
   Aqui temos um programa que lança processos concorrentes e espera o término dos mesmos.
  
   EXERCICIOS:
      1) modifique o número de processos e o numero de iteracoes de cada processo
      2) avalie o resultado obtido do ponto de vista da velocidade relativa entre os processos e da justiça.
	   R: Todos os processos são executados (há justiça), em ordem aleatória
      3) observe que os itens comunicados pelo canal fin são vazios.
         isto significa que o importante neste caso é somente a sincronização.
      4) 'fin' é um canal sincrono.
         faria diferença se 'fin' fosse assíncrono, ou seja, se tivesse um buffer para armazenar itens ?
	   R: Sim, pois com canais assíncronos, não leitura e escrita não são bloqueadas
*/
package main

import (
	"fmt"
)

const chNumber int = 10

func algoConcorrente(id int, par int, fin chan struct{}) {
	for i := 0; i < par; i++ {
		fmt.Printf("Canal(%d), iteracao(%d)\n", id, i)
	}
	fin <- struct{}{} // sinaliza final
}

func main() {
	fin := make(chan struct{})

	// cria 5 rotinas concorrentes
	for i := 0; i < chNumber; i++ {
		go algoConcorrente(i, 3, fin) // passa canal fin para avisar o termino
	}

	// espera o termino das rotinas
	for i := 0; i < chNumber; i++ {
		<-fin // wait for 5 processes to write in ch
	}
	fmt.Println("fim")
}