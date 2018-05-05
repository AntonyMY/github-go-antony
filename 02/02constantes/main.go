package main

import (
	"fmt"
)

func main() {
	//const nombre = "Pedro"
	//fmt.Println(nombre)

	var a int
	var b int8
	var numero int

	n := "José"
	p := "Sifuentes"

	a = 121212
	b = 5

	//casting
	c := a + int(b)
	fmt.Println(c)
	fmt.Printf("Hola %s %s cómo estás?\n", n, p)
	fmt.Printf("c es un tipo de dato %T\n", c)
	fmt.Println(numero)

	//prioridad arimética

}
