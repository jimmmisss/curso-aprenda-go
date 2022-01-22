package main

import "fmt"

var y = "Olá brow"

var a int
var b float64
var c float32
var d bool
var e string

func main() {
	helloWorld()
}

func helloWorld() {
	/*fmt.Println("### HELLO WORLD ###")
	numeroDeCatacteres, err := fmt.Println("Hello world", "Wesley", 666)
	fmt.Println(numeroDeCatacteres, err)
	variaveis()*/
	tipos()
}

func variaveis() {
	fmt.Println("### VARIÁVEIS ###")

	x := 10.01

	fmt.Printf("X: %v, %T\n", x, x)
	fmt.Printf("Y: %v, %T\n\n", y, y)

	x, z := 80.06, "Isadora"
	fmt.Println(x, z)
}

func tipos() {
	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)
	fmt.Printf("d: %v, %T\n", d, d)
	fmt.Printf("e: %v, %T\n", e, e)
}
