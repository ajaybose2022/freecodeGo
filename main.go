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

const (
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

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

	st := "This is a string"
	by := []byte(st)
	fmt.Printf("%v, %T\n", st, st)
	fmt.Printf("%v, %T\n", by, by)

	// Rune

	ru := 'a'
	fmt.Printf("%v, %T\n", ru, ru)

	// Constants

	fileSize := 412345707080930023670.

	fmt.Printf("%.2fKB \n", fileSize/KB)
	fmt.Printf("%.2fMB \n", fileSize/MB)
	fmt.Printf("%.2fGB \n", fileSize/GB)
	fmt.Printf("%.2fTB \n", fileSize/TB)
	fmt.Printf("%.2fPB \n", fileSize/PB)
	fmt.Printf("%.2fEB \n", fileSize/EB)
	fmt.Printf("%.2fZB \n", fileSize/ZB)
	fmt.Printf("%.2fYB \n", fileSize/YB)

	// Arrays

	grades := [3]uint8{97, 85, 93}
	grades1 := [...]uint8{97, 85, 93, 88} // ... specifies, any size
	fmt.Printf("Grades  : %v\n", grades)
	fmt.Printf("Grades1 : %v\n", grades1)

	var students [3]string
	fmt.Printf("Students : %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students : %v\n", students)
	fmt.Printf("Students #1: %v\n", students[1])
	fmt.Printf("Number of Students : %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(identityMatrix)

	var array1 = [...]int{1, 2, 3}
	var array2 = array1
	var array3 = &array1
	array2[1] = 8
	fmt.Printf("array1 : %v\n", array1)
	fmt.Printf("array2 : %v\n", array2)
	fmt.Printf("array3 : %v\n", array3)
	array3[2] = 76
	fmt.Printf("array1 : %v\n", array1)
	fmt.Printf("array2 : %v\n", array2)
	fmt.Printf("array3 : %v\n", array3)
}
