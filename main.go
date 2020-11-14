package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//3 10 2001
	//12 11 2020
	//9 1 19
	fmt.Println("ingresa fu fecha de nacimiento ejmplo 3/10/2001")
	scanner.Scan()
	fechaN := scanner.Text()
	//Declaracion de arrays
	bytes := []byte(fechaN)
	bucleArray := []int{}
	datosManejables := []int{}
	for index := range bytes {
		index := index
		bucleArray = append(bucleArray, int(bytes[index]))
		if bucleArray[index] != 47 {
			datosManejables = append(datosManejables, int(bytes[index]))
		}

	}

	nValidos := [10]int{49, 50, 51, 52, 53, 54, 55, 56, 57, 48}

	for index := range datosManejables {
		i := index

		for j := range nValidos {
			if datosManejables[i] == nValidos[j] {

			}

		}

	}

}
