package multi

import (
	"fmt"
	"strconv"
	"strings"
)

//Multiplicar multiplica el resultado
func Multiplicar(operacion string) int {
	valor := strings.Split(operacion, "*")
	resultado := 0

	for i := 0; i < len(valor); i++ {
		num, err := strconv.Atoi(valor[i])

		if err != nil {
			fmt.Println(err)
			fmt.Println("Ingrese un numero valido")
		} else {
			if resultado == 0 {
				resultado = num
			} else {
				resultado *= num
			}
		}
	}
	return resultado
}
