package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
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

	switch {
	case numeros[0] != "0":
		aux := string(numeros[0]) + string(numeros[1])

		dia, _ = strconv.Atoi(aux)

	case numeros[0] == "0":
		dia, _ = strconv.Atoi(numeros[1])

	}

	mes := 0

	switch {
	case numeros[2] == "0":
		mes, _ = strconv.Atoi(numeros[3])
	case numeros[2] != "0":
		aux := string(numeros[2]) + string(numeros[3])
		mes, _ = strconv.Atoi(aux)

	}

	yearString := numeros[len(numeros)-4] + numeros[len(numeros)-3] + numeros[len(numeros)-2] + numeros[len(numeros)-1]
	yearInt, _ := strconv.Atoi(yearString)

	//fecha actual
	year, _, d := time.Now().Date()

	diasMes := mesIntFunc() - diasMes(mes)
	diasYear := diasYear(yearInt, year)
	//me falta restar el mes y el dia actual

	diasDeVida := (diasYear + diasMes + d - dia)

	fmt.Println("Tienes", diasDeVida, "dias de vida")

}

//esta funcion va a devolver la suma de los dias de los meses dependiendo
func mesIntFunc() int {
	enero := 31
	febrero := 28
	marzo := 31
	abril := 30
	mayo := 31
	junio := 30
	julio := 31
	agosto := 31
	septiembre := 30
	octubre := 31
	noviembre := 30
	diciembre := 31
	_, mes, _ := time.Now().Date()
	switch {
	case mes == time.January:
		return enero
	case mes == time.February:
		return enero + febrero
	case mes == time.March:
		return enero + febrero + marzo
	case mes == time.April:
		return enero + febrero + marzo + abril
	case mes == time.May:
		return enero + febrero + marzo + abril + mayo
	case mes == time.June:
		return enero + febrero + marzo + abril + mayo + junio
	case mes == time.July:
		return enero + febrero + marzo + abril + mayo + junio + julio
	case mes == time.August:
		return enero + febrero + marzo + abril + mayo + junio + julio + agosto
	case mes == time.September:
		return enero + febrero + marzo + abril + mayo + junio + julio + agosto + septiembre
	case mes == time.October:
		return enero + febrero + marzo + abril + mayo + junio + julio + agosto + septiembre + octubre
	case mes == time.November:
		return enero + febrero + marzo + abril + mayo + junio + julio + agosto + septiembre + octubre + noviembre
	case mes == time.December:
		return enero + febrero + marzo + abril + mayo + junio + julio + agosto + septiembre + octubre + noviembre + diciembre
	}
	return 0
}

//esta funcion devuelve un nuemro dependiendo en mes de nacimiento ejemplo octubre es igual a 31+(todos los meces anteriores)
//osea regresa la suma de los dias corespondientes al mes
func diasMes(mes int) int {
	enero := 31
	febrero := 28
	marzo := 31
	abril := 30
	mayo := 31
	junio := 30
	julio := 31
	agosto := 31
	septiembre := 30
	octubre := 31
	noviembre := 30
	diciembre := 31
	switch {
	case mes == 1:
		return enero
	case mes == 2:
		return enero + febrero
	case mes == 3:
		return enero + febrero + marzo
	case mes == 4:
		return enero + febrero + marzo + abril
	case mes == 5:
		return enero + febrero + marzo + abril + mayo
	case mes == 6:
		return enero + febrero + marzo + abril + mayo + junio
	case mes == 7:
		return enero + febrero + marzo + abril + mayo + junio + julio
	case mes == 8:
		return enero + febrero + marzo + abril + mayo + junio + julio + agosto
	case mes == 9:
		return enero + febrero + marzo + abril + mayo + junio + julio + agosto + septiembre
	case mes == 10:
		return enero + febrero + marzo + abril + mayo + junio + julio + agosto + septiembre + octubre
	case mes == 11:
		return enero + febrero + marzo + abril + mayo + junio + julio + agosto + septiembre + octubre + noviembre
	case mes == 12:
		return enero + febrero + marzo + abril + mayo + junio + julio + agosto + septiembre + octubre + noviembre + diciembre

	}

	return 0

}

//esta funcion devuelve la resta entre el no de naciminto y el ano actual
func diasYear(yearInt int, year int) int {

	return (year - yearInt) * 365
}

//version1

