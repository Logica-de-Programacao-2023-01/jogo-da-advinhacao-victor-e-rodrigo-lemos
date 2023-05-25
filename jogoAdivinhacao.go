package main

import (
	"fmt"
	"math/rand"
)

func Random() int {

	return rand.Intn(100) + 1
}
func main() {
	num := Random()
	repetir := true
	var numUser int
	tentativas := []int{}
	contador := 0
	var simNao string

	for repetir {
		fmt.Scanln(&numUser)
		if numUser > num {
			fmt.Println("O número é menor que", numUser)
			contador++
			continue
		} else if numUser < num {
			fmt.Println("O número é maior que", numUser)
			contador++
			continue
		} else {
			fmt.Println("Parabéns, você acertou!")
			contador++
			tentativas = append(tentativas, contador)
			contador = 0
		}
		fmt.Println("Você deseja jogar novamente? (s/n): ")
		fmt.Scan(&simNao)
		if simNao == "n" {
			repetir = false
		} else {
			num = Random()
			fmt.Scanln(&numUser)
		}
	}
	for i := 0; i < len(tentativas); i++ {
		fmt.Printf("Você utilizou %d na %dª tentativa\n", tentativas[i], i+1)
	}
}
