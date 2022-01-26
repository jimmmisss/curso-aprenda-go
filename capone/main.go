package main

import (
	"fmt"
)

const ui int = 1000

var y = "Olá brow"

var a int
var b float64
var c float32
var d bool
var e string

func main() {
	loopFor()
}

func helloWorld() {
	fmt.Println("### HELLO WORLD ###")
	numeroDeCatacteres, err := fmt.Println("Hello world", "Wesley", 666)
	fmt.Println(numeroDeCatacteres, err)
	variaveis()
	tipos()
	exercicio01()
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

func exercicio01() {
	x := 42
	y := "Lemmy Kilmster"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func decimalBinarioEHexadecimal() {
	x := 100
	fmt.Printf("%d, %#x, %b", x, x, x)
}

func expressoes() {
	a := (10 == 10)
	b := (10 != 10)
	c := (10 <= 10)
	d := (10 < 10)
	e := (10 >= 10)
	f := (10 > 10)

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}

func loopFor() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("%d - %v\n", x, string(x))
	}
}
