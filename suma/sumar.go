package suma

import (
	"fmt"
	"strconv"
	"strings"
)

func Sumar(operacion string) int {
	valor := strings.Split(operacion, "+")
	resultado := 0

	for i := 0; i < len(valor); i++ {
		num, err := strconv.Atoi(valor[i])

		if err != nil {
			fmt.Println(err)
			fmt.Println("Ingrese un numero valido")
		} else {
			resultado += num
		}

	}
	return resultado
}
