package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CPabb-27/Calculadora-Backend/div"
	"github.com/CPabb-27/Calculadora-Backend/multi"
	"github.com/CPabb-27/Calculadora-Backend/resta"
	"github.com/CPabb-27/Calculadora-Backend/suma"
)

func main() {
	fmt.Println("Calculadora Backend")

	resultados()

}

func resultados() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	resultado := 0
	if strings.Contains(operacion, "+") {
		resultado = suma.Sumar(operacion)
	} else if strings.Contains(operacion, "-") {
		resultado = resta.Restar(operacion)
	} else if strings.Contains(operacion, "*") {
		resultado = multi.Multiplicar(operacion)
	} else if strings.Contains(operacion, "/") {
		resultado = div.Division(operacion)
	} else {
		fmt.Printf("ELIJA UN OPERADOR VALIDO\n")
	}
	fmt.Println("El resultado es: ", resultado)

}
