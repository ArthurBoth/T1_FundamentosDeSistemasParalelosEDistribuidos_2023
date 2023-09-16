/* por Arthur Antunes de Souza Both e Gabriel Moraes Ferreira

     1) A operação de busca de um elemento v, dizendo true se encontrou v na árvore, ou false
			R: Linhas [122 - 130]
     2) A operação de busca concorrente de um elemento, que informa imediatamente
          por um canal se encontrou o elemento (sem acabar a busca), ou informa
          que nao encontrou ao final da busca
		    R: Linhas [133 - 159]
     3) A operação que escreve todos pares em um canal de saidaPares e
          todos impares em um canal saidaImpares, e ao final avisa que acabou em um canal fin
            R: Linhas [162 - 176]
     4) A versao concorrente da operação acima, ou seja, os varios nodos sao testados
          concorrentemente se pares ou impares, escrevendo o valor no canal adequado
		    R: Linhas [178 - 210]
*/
package main

import (
	"fmt"
)

type Nodo struct {
	v int
	e *Nodo
	d *Nodo
}

func caminhaERD(r *Nodo) {
	if r != nil {
		caminhaERD(r.e)
		fmt.Print(r.v, ", ")
		caminhaERD(r.d)
	}
}

// -------- SOMA ----------
// soma sequencial recursiva
func soma(r *Nodo) int {
	if r != nil {
		//fmt.Print(r.v, ", ")
		return r.v + soma(r.e) + soma(r.d)
	}
	return 0
}

// funcao "wraper" retorna valor
// internamente dispara recursao com somaConcCh
// usando canais
func somaConc(r *Nodo) int {
	s := make(chan int)
	go somaConcCh(r, s)
	return <-s
}
func somaConcCh(r *Nodo, s chan int) {
	if r != nil {
		s1 := make(chan int)
		go somaConcCh(r.e, s1)
		go somaConcCh(r.d, s1)
		s <- (r.v + <-s1 + <-s1)
	} else {
		s <- 0
	}
}

// ---------   agora vamos criar a arvore e usar as funcoes acima

func main() {
	root := &Nodo{v: 10,
		e: &Nodo{v: 5,
			e: &Nodo{v: 3,
				e: &Nodo{v: 1, e: nil, d: nil},
				d: &Nodo{v: 4, e: nil, d: nil}},
			d: &Nodo{v: 7,
				e: &Nodo{v: 6, e: nil, d: nil},
				d: &Nodo{v: 8, e: nil, d: nil}}},
		d: &Nodo{v: 15,
			e: &Nodo{v: 13,
				e: &Nodo{v: 12, e: nil, d: nil},
				d: &Nodo{v: 14, e: nil, d: nil}},
			d: &Nodo{v: 18,
				e: &Nodo{v: 17, e: nil, d: nil},
				d: &Nodo{v: 19, e: nil, d: nil}}}}

	fmt.Println()
	fmt.Print("Valores na árvore: ")
	caminhaERD(root)
	fmt.Println()
	fmt.Println()

	fmt.Println("Soma: ", soma(root))
	fmt.Println("SomaConc: ", somaConc(root))
	fmt.Println()
	
	fmt.Println("Busca 7: ", busca(root, 7))
	fmt.Println("BuscaConc 7: ", buscaConc(root, 7))
	fmt.Println("Busca 2: ", busca(root, 2))
	fmt.Println("BuscaConc 2: ", buscaConc(root, 2))
	fmt.Println()
	
	pares := make([]int, 0)
	impares := make([]int, 0)
	retornaParImparRec(root, &pares, &impares)
	fmt.Println("Números pares na árvore   (recursivo): ", pares)
	fmt.Println("Números impares na árvore (recursivo): ", impares)
	fmt.Println()
	
	pares = []int{} // esvazia os slices
	impares = []int{}
	retornaParImparConc(root, &pares, &impares)
	fmt.Println("Números pares na árvore   (concorrente): ", pares)
	fmt.Println("Números impares na árvore (concorrente): ", impares)
	fmt.Println()
}
//---------------------------------------------------------------------------------------------//
// Nossas implementações

// Resposta 1 linhas [122 - 130]
func busca(r* Nodo, x int) bool {
	if r != nil {
		if r.v == x {
			return true
		}
		return busca(r.e, x) || busca(r.d, x)
	}
	return false
}

// Resposta 2 linhas [133 - 159]
func buscaConc(r *Nodo, x int) bool {
	b := make(chan struct{})
	fim := make(chan struct{})
	go buscaConcCh(r, x, b, fim)
	for{
		select {
		case <- fim:
			return false
		case <- b:
			return true
		}
	}
}

func buscaConcCh(r *Nodo, x int, b chan struct{}, fim chan struct{}) {
	end := make(chan struct{} ,2)
	if r != nil {
		if r.v == x {
			b <- struct{}{}
		}
		go buscaConcCh(r.e, x, b, end)
		go buscaConcCh(r.d, x, b, end)
		<- end
		<- end
	}
	fim <- struct{}{}
}

// Resposta 3 linhas [162 - 176]
func retornaParImparRec(r *Nodo, pares *[]int, impares *[]int){
	retornaParImpar(r, pares, impares)
}

func retornaParImpar(r *Nodo, saidaP *[]int, saidaI *[]int) {
	if r != nil {
		retornaParImpar(r.e, saidaP, saidaI)
		if r.v % 2 == 0 {
			*saidaP = append(*saidaP, r.v)
		} else {
			*saidaI = append(*saidaI, r.v)
		}
		retornaParImpar(r.d, saidaP, saidaI)
	}
}

// Resposta 4 linhas [178 - 210]
func retornaParImparConc(r *Nodo, pares *[]int, impares *[]int){
	saidaP := make(chan int)
	saidaI := make(chan int)
	fin := make(chan struct{})
	go retornaParImparConcCh(r, saidaP, saidaI, fin)
	for  {
		select {
		case p := <- saidaP:
			*pares = append(*pares, p)
		case i := <- saidaI:
			*impares = append(*impares, i)
		case <- fin:
			return
		}
	}
}

func retornaParImparConcCh(r *Nodo, saidaP chan int, saidaI chan int, fin chan struct{}) {
	fim := make(chan struct{} ,2)
	if r != nil {
		if r.v % 2 == 0 {
			saidaP <- r.v
		} else {
			saidaI <- r.v
		}
		go retornaParImparConcCh(r.e, saidaP, saidaI, fim)
		go retornaParImparConcCh(r.d, saidaP, saidaI, fim)
		<- fim
		<- fim
	}
	fin <- struct{}{}
}