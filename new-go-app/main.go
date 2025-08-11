package main

import (
	"fmt"
)

func main()
	var intNum int = 32767
	intNum = intNum + 1 
	fmt.printLn(intNum)

	var floatNum float64 = 123123123.9
	fmt.printLn(floatNum)

	var floatNum32 float32 = 10.1 
	var intNum32 int32 = 2 
	var result float32 = floatNum32 * float32(intNum32)
	fmt.Println("Result of float32 * int32:", result)

	var intNum1 int = 3 
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("y"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false 
	fmt.Println(myBoolean)

	var intNum3 rune 
	fmt.Println(intNum3)

	var myVar string = foo()
	fmt.Println(myVar)

	var1, var2 := 1, 2 
	fmt.Println(var1, var2)

	const myConst string = "const value"
	fmt.Println(myConst)
}
