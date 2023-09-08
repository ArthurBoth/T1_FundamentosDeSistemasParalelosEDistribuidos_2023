/* por Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
 EXERCÍCIO:  dado o programa abaixo
    1) quantos processos concorrentes são gerados ?
		R: 41, um para cada chamada da 'funcaoA' (definida em N) e um para o 'main'
    2) execute e observe: que se pode supor sobre a velocidade relativa dos mesmos ?
		R: Cada processo executa em uma velocidade própria inpependente de outros processos,
			 mas todos ao mesmo tempo
*/

package main

import (
	"fmt"
	"time"
)

var N int = 40

func funcaoA(id int, s string) {
	for {
		fmt.Println(s, id)
	}
}

func geraNespacos(n int) string {
	s := "  "
	for j := 0; j < n; j++ {
		s = s + "   "
	}
	return s
}

func main() {
	for i := 0; i < N; i++ {
		go funcaoA(i, geraNespacos(i))
	}
	for true {
		time.Sleep(100 * time.Millisecond)
	}
}
