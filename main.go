package main

import (
	"fmt"
	"strconv"
)

/*
	Group the vars together

var (

	actorName    string = "Elizabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11

)
*/
func main() {
	var i int = 42
	var j float32
	var k string
	//var k int -- Not recommended
	//k = 67
	fmt.Printf("%v %T\n", i, i)
	j = float32(i)
	fmt.Printf("%v %T\n", j, j)
	k = strconv.Itoa(i)
	fmt.Printf("%v %T\n", k, k)
	//fmt.Println(k)

	var n bool = false // Result of a logical operation
	fmt.Printf("%v, %T\n", n, n)

	// int and int8 are mismatched
	// we need to explicitly type cast every mismatch

	fmt.Println("Math Operations")
	a := 10
	b := 3
	fmt.Println(a + b) // add
	fmt.Println(a - b) // subtract
	fmt.Println(a * b) // multiply
	fmt.Println(a / b) // divide
	fmt.Println(a % b) // remainder

	fmt.Println("Logical Operations")

	fmt.Println(a & b)  // logical and
	fmt.Println(a | b)  // logical or
	fmt.Println(a ^ b)  // logical not
	fmt.Println(a &^ b) // logical and not

	fmt.Println("Bit Operations")

	fmt.Println(a << 2) // Shift left (Means Multiply by 2 to power of number )  == 10 * 2 ^ 2 = 10 * 4 = 40
	fmt.Println(a >> 2) // Shift Right (Means Divide by 2 to power of number )  == 10 /( 2 ^ 2 ) = 10 / 4 = 2 ( since it is int and not float)

	// Float
	var fl float64 = 3.14
	fmt.Printf("%v, %T \n", fl, fl)
	fl = 13.7e72
	fmt.Printf("%v, %T \n", fl, fl)
	fl = 2.1e14
	fmt.Printf("%v, %T \n", fl, fl)

	// Complex Numbers
	fmt.Println("Complex Numbers")

	complex1 := 1 + 2i
	complex2 := 4 + 5.3i

	fmt.Print(complex1 + complex2)
	fmt.Printf("\tReal Part: %v, Imaginary Part: %v\n", real(complex1+complex2), imag(complex1+complex2))
	fmt.Print(complex1 - complex2)
	fmt.Printf("\tReal Part: %v, Imaginary Part: %v\n", real(complex1-complex2), imag(complex1-complex2))
	fmt.Print(complex1 * complex2)
	fmt.Printf("\tReal Part: %v, Imaginary Part: %v\n", real(complex1*complex2), imag(complex1*complex2))
	fmt.Print(complex1 / complex2)
	fmt.Printf("\tReal Part: %v, Imaginary Part: %v\n", real(complex1/complex2), imag(complex1/complex2))

	fmt.Println("Print Real and Imag Seperately")
	fmt.Printf("Real Part: %v, Imaginary Part: %v\n", real(complex1+complex2), imag(complex1+complex2))

	// Strings

	s := "This is a string"
	by := []byte(s)
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", by, by)

}
