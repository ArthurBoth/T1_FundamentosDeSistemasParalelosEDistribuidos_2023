/* por Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
  
   ASSUNTO - Compreensão de concorrência e canais com buffer
  
   EXERCÍCIO:
       1) Avalie o comportamento do programa para tamBuff
          0 e 10.   Voce consegue explicar a diferença ?
		  R: Com um buffer de tamanho 0, o programa precisa esperar para escrever no canal
		     quando ele está cheio, enquanto que com um buffer de tamanho 10, o programa
			 pode eontinuar escrevendo no canal até que ele esteja cheio.
			 Na prática, o programa com buffer 0 imprime a escrita e a leitura de cada valor
			 alternadamente, enquanto com buffer 10, é mais aleatório.
       2) Qual versao tem maior nivel de concorrencia ?
		  R: A com buffer 10, pois ambos os processos podem escrever e ler do canal ao mesmo tempo,
		     enquanto no de buffer 0, um processo precisa esperar o outro terminar para poder escrever.
       3) Faça uma versão que tem vários processos destino
          que podem consumir os dados de forma não determinística.
          Ou seja, processos diferentes podem consumir quantidades
          diferentes de itens,  conforme sua velocidade.
          Como você coordenaria o término dos processos depois do
          consumo dos N valores ?
*/
package main

const N = 100
const M = 25
const tamBuff = 0

func fonteDeDados(saida chan int) {
	for i := 1; i < N; i++ {
		println(i, " -> ")
		saida <- i
	}
}

func destinoDosDados(entrada chan int, fim chan struct{}) {
	for i := 1; i < M; i++ {
		v := <-entrada
		println("                  -> ", v)
	}
	fim <- struct{}{}
}

func end(fim chan struct{}){
	for i:=0; i<(N/M); i++{
		<-fim
	}
}

func main() {
	c := make(chan int, tamBuff)
	d := make(chan struct{})
	go fonteDeDados(c)
	go destinoDosDados(c,d)
	go destinoDosDados(c,d)
	go destinoDosDados(c,d)
	go destinoDosDados(c,d)
	end(d)
}
