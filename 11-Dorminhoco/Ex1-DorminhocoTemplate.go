/* por Arthur Antunes de Souza Both e Gabriel Moraes Ferreira

     Um template para criar um anel generico.
     Adapte para o problema do dorminhoco.
     Nada está dito sobre como funciona a ordem de processos que batem.
     O ultimo leva a rolhada ...
     ESTE  PROGRAMA NAO FUNCIONA.    É UM RASCUNHO COM DICAS.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var fin = make(chan struct{})

const NJ = 5           // numero de jogadores
const M = 4            // numero de cartas

type carta string      // carta é um string

var ch [NJ]chan carta  // NJ canais de itens tipo carta  

var bateu = make(chan struct{}, NJ) // avisa que bateu
var countEnds = 1 // contador de jogadores que já bateram (inicia em 1)
var createPlayers = make(chan struct{}, NJ)

func jogador(id int, in chan carta, out chan carta, cartasIniciais []carta) {
	createPlayers <- struct{}{}
	mao := cartasIniciais
    var cartaRecebida carta
	var cartaEnviada carta
	

	fmt.Printf("[%d] Mão inicial(jogador): %v\n",id,mao)
	for {
		select {
		case cartaRecebida = <- in:
			fmt.Printf("[%d] Mão mais antes(jogador): %v\n",id,mao)
			fmt.Printf("[%d] recebeu \"%s\"\n", id, cartaRecebida)

			cartaEnviada = choice(mao)
			fmt.Printf("[%d] Mão antes(jogador): %v\n",id,mao)
			mao = append(mao, cartaRecebida)
			fmt.Printf("[%d] Mão meio(jogador): %v\n",id,mao)
			mao = remove(mao, cartaEnviada)
			fmt.Printf("[%d] Mão depois(jogador): %v\n",id,mao)
			
			fmt.Printf("[%d] enviou \"%s\" para [%d]\n", id, cartaEnviada, (id+1)%NJ)
			out <- cartaEnviada
		case <- bateu:
			bater(id)
		default:/*
			if checkEnd(mao) {
				bater(id)
			}*/
		}
	}
}

func main() {
	baralho := makeDeck()
	fmt.Println("baralho: ",baralho)
	
	for i := 0; i < NJ; i++ {
		ch[i] = make(chan carta)
	}
	
	for i := 0; i < NJ; i++ {
		cartasEscolhidas := baralho[:M]
		baralho = baralho[M:]
		fmt.Printf("[%d] Mão inicial(main): %v\n",i ,cartasEscolhidas)
		go jogador(i, ch[i], ch[(i+1)%NJ], cartasEscolhidas)
	}
		fmt.Println("\n")

	for i := 0; i < NJ; i++ {
		<- createPlayers
	}
	ch[0] <- baralho[0]
	time.Sleep(50 * time.Millisecond)
	//fim()
}

//--------------------------------------------------------
func choice(mao []carta) carta {
	if (contains(mao, "Joker")){
		return "Joker"
	}
	var c int = rand.Intn(len(mao))
	return mao[c]
}

func index(mao []carta, c carta) int {
	for i, n := range mao {
		if n == c {
			return i
		}
	}
	return -1
}

func contains(mao []carta, x carta) bool {
	for _, n := range mao {
		if x == n {
			return true
		}
	}
	return false
}

func remove(mao []carta, c carta) []carta {
	i := index(mao, c)
    mao[i] = mao[len(mao)-1]
    return mao[:len(mao)-1]
}

func bater(id int) {
	fmt.Printf("[%d] bateu\n", id)
	for i := 0; i < (NJ-countEnds); i++ {
		bateu <- struct{}{}
	}
	countEnds++
	fin <- struct{}{}
}

func checkEnd(mao []carta) bool {
	c := mao[0]
	for _, n := range mao {
		if c != n {
			return false
		}
	}
	return true
}

func makeDeck() []carta {
	var deck []carta
	for i := 0; i < NJ; i++ {
		for j := 0; j < M; j++ {
			deck = append(deck, carta(fmt.Sprintf("%d",i)))
		}
	}
	deck = append(deck, "Joker") // coringa
	deck = shuffle(deck)
	return deck
}

func shuffle(deck []carta) []carta{
	for i := range deck { // embaralha
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func fim() {
	for i := 0; i < NJ; i++ {
		<- fin
	}
}