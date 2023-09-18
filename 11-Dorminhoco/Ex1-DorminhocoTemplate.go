/* por Arthur Antunes de Souza Both e Gabriel Moraes Ferreira */

package main

import (
	"fmt"
	"math/rand"
)

var fin = make(chan int)

const NJ = 5           // numero de jogadores
const M = 4            // numero de cartas

type carta string      // carta é um string

var ch [NJ]chan carta  // NJ canais de itens tipo carta  

var bateu = make(chan struct{}, 1) // avisa que bateu
var createPlayers = make(chan struct{}, NJ)

func jogador(id int, in chan carta, out chan carta, cartasIniciais []carta) {
	createPlayers <- struct{}{}
	mao := cartasIniciais
	var cartaEnviada carta
	var cartaRecebida carta
	var aux []carta = make([]carta, 0) // mão auxiliar que existe por causa de ponteiros de slice
	

	for {
		select {
		case cartaRecebida = <- in:
			
			cartaEnviada = mao[choice(mao)]
			aux = append(mao, cartaRecebida)
			mao = remove(aux, cartaEnviada)
			
		case <- bateu:
			go bater(id)
			return
		default:
			if (cartaEnviada != ""){
				fmt.Printf("[%d] enviou \033[36m%s\033[0m para [%d]\n", id, cartaEnviada, (id+1)%NJ)
				out <- cartaEnviada
				cartaEnviada = ""
			}
			if checkEnd(mao) {
				go bater(id)
				return
				}
		}
	}
}

func main() {
	baralho := makeDeck()
	fmt.Println("baralho: ",baralho)
	
	for i := 0; i < NJ; i++ {
		ch[i] = make(chan carta, 1)
	}
	
	for i := 0; i < NJ; i++ {
		cartasEscolhidas := baralho[:M]
		baralho = baralho[M:]
		fmt.Printf("[%d] Mão inicial(main): %v\n",i ,cartasEscolhidas)
		go jogador(i, ch[i], ch[(i+1)%NJ], cartasEscolhidas)
	}
	
	for i := 0; i < NJ; i++ {
		<- createPlayers
	}

	fmt.Println("\n")

	ch[0] <- baralho[0]

	fim()
}

//--------------------------------------------------------
func choice(mao []carta) int {
	jokerIndex := index(mao, "@")
	if (jokerIndex >= 0){
		return jokerIndex
	}

	return rand.Intn(len(mao))
}

func index(mao []carta, c carta) int {
	for i, n := range mao {
		if n == c {
			return i
		}
	}
	return -1
}

func remove(mao []carta, c carta) []carta {
	ret := make([]carta, 0)
	aux := make([]carta, 0)
	i := index(mao, c)

	aux = append(aux, mao...)

    aux[i] = mao[len(mao)-1]
	ret = append(ret, aux[:len(aux)-1]...)
    return ret
}

func bater(id int) {
	bateu <- struct{}{}
	fin <- id
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
	deck = append(deck, "@") // coringa
	deck = shuffle(deck)
	return deck
}

func shuffle(deck []carta) []carta{
	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func fim() {
	var players []int

	for i := 0; i < NJ; i++ {
		players = append(players, i)
	}

	for i := 0; i < NJ-1; i++ {
		id := <- fin
		fmt.Printf("\033[33m[%d] bateu\n\033[0m", id)
		players = removeIndex(players, id)
	}

	fmt.Printf("\033[31m%v é o dorminhoco\n\033[0m", players)
}

func removeIndex(slice []int, id int) []int {
	index := findIndex(slice, id)

	ret := make([]int, 0)
	ret = append(ret, slice[:index]...)
	ret = append(ret, slice[index+1:]...)

	return ret
}

func findIndex(slice []int, index int) int {
	for i, n := range slice {
		if n == index {
			return i
		}
	}
	return -1
}