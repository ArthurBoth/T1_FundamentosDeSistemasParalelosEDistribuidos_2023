/* por Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
 EXERCÍCIO:  dado o programa abaixo
     considerando-se como estados os valores da tripla x,y,z
     qual o diagrama de estados e transicoes que representa
     1) a questaoStSp2()  ?
	 R: 				x,y
	 (nível 0)			0,0
	 (nível 1)		1,0		0,1
	 (nível 2)	2,0		1,1		0,2
	 (nível 3)		2,1		1,2
	 (nível 4)			2,2
     2) a questaoStSp3()  ?
	 R: 		x,y,z (não é possível representar este diagrama num espaço bidimensional, mas segue a lista de estados)
	 (nível 0): 0,0,0
	 (nível 1): 1,0,0 ; 0,1,0 ; 0,0,1
	 (nível 2): 2,0,0 ; 1,1,0 ; 1,0,1 ; 0,2,0 ; 0,1,1 ; 0,0,2
	 (nível 3): 2,0,1 ; 2,1,0 ; 1,2,0 ; 0,2,1 ; 0,1,2 ; 1,0,2
	 (nível 4): 2,0,2 ; 2,1,1 ; 2,2,0 ; 1,2,1 ; 0,2,2 ; 1,1,2
	 (nível 5): 2,1,2 ; 2,2,1 ; 1,2,2
	 (nível 6): 2,2,2
*/

package main

//---------------------------

var x, y, z int = 0, 0, 0

func px() {
	x = 1
	x = 2
}

func py() {
	y = 1
	y = 2
}

func pz() {
	z = 1
	z = 2
}

func questaoStSp2() {
	go px()
	py()
	for {
	}
}

func questaoStSp3() {
	go px()
	go py()
	pz()
	for {
	}
}

func main() {
	questaoStSp2()
	questaoStSp3()
}
