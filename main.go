package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	datosReales := []int{}

	for index := range datosManejables {
		i := index

		for j := range nValidos {

			if datosManejables[i] == nValidos[j] {
				datosReales = append(datosReales, datosManejables[i])

			}

		}

	}
	numeros := []string{}
	for index := range datosReales {
		switch datosReales[index] {
		case 49:
			numeros = append(numeros, "1")
		case 50:
			numeros = append(numeros, "2")
		case 51:
			numeros = append(numeros, "3")
		case 52:
			numeros = append(numeros, "4")
		case 53:
			numeros = append(numeros, "5")
		case 54:
			numeros = append(numeros, "6")
		case 55:
			numeros = append(numeros, "7")
		case 56:
			numeros = append(numeros, "8")
		case 57:
			numeros = append(numeros, "9")
		case 48:
			numeros = append(numeros, "0")

		}
	}
	dia := 0
	fmt.Println(numeros)
	switch {
	case numeros[0] != "0":
		aux := string(numeros[0]) + string(numeros[1])

		dia, _ = strconv.Atoi(aux)

	case numeros[0] == "0":
		dia, _ = strconv.Atoi(numeros[1])

	}
	fmt.Println("dia ", dia)

	mes := 0
	switch {
	case numeros[2] == "0":
		mes, _ = strconv.Atoi(numeros[3])
	case numeros[2] != "0":
		aux := string(numeros[2]) + string(numeros[3])
		mes, _ = strconv.Atoi(aux)

	}
	fmt.Println("mes ", mes)

}
