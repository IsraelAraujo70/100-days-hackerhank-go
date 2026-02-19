package main

import "fmt"

const Constante string = "Public variable"
const constante = "private variable"

func main() {
	var name string = "Israel"
	fmt.Println(name)
	fmt.Printf("Variavel é do tipo: %T \n", name)

	var isBool bool = true
	fmt.Println(isBool)
	fmt.Printf("Variavel é do tipo: %T \n", isBool)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloatVal uint8 = 255
	fmt.Println(smallFloatVal)
	fmt.Printf("Variable is of type: %T \n", smallFloatVal)

	// default values and aliases
	var emptyVariableInt int
	fmt.Println(emptyVariableInt)
	fmt.Printf("Variable is of type: %T \n", emptyVariableInt)

	//no var
	number := 3000
	fmt.Println(number)

	fmt.Println(Constante)
}
