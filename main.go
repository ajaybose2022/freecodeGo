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
}
