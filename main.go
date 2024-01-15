package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"sync"
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

type Doctor struct {
	number     uint8
	actorName  string
	episodes   []string
	companions []string
}

//Inheritance is not present, only use embedded

type Animal struct {
	Name   string `validate:"required max:100"` //`json:"gender" validate:"required,gender_custom_validation"`
	Origin string
}
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

// WaitGroup
var wg = sync.WaitGroup{}

// Mutex
var mutex = sync.RWMutex{}

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

	// Slices

	var slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := slice1[:]
	slice3 := slice1[3:]
	slice4 := slice1[:6]
	slice5 := slice1[3:6]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Printf("Length : %v\n", len(slice1))
	fmt.Printf("Capacity : %v\n", cap(slice1))

	intSlice := make([]int, 3, 20)
	fmt.Println(intSlice)
	fmt.Printf("Length : %v\n", len(intSlice))
	fmt.Printf("Capacity : %v\n", cap(intSlice))
	intSlice = append(intSlice, 1)
	fmt.Printf("Length : %v\n", len(intSlice))
	fmt.Printf("Capacity : %v\n", cap(intSlice))

	//maps - slice can never be a type of map

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        2786596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	fmt.Println(statePopulations)

	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)
	_, isIllinoisPresent := statePopulations["Illinois"]
	fmt.Println(isIllinoisPresent)
	delete(statePopulations, "Illinois")
	fmt.Println(statePopulations)
	_, isIllinoisPresent = statePopulations["Illinois"]
	fmt.Println(isIllinoisPresent)

	// Struct can use field name as well as positional, but recommended to use named field
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		episodes: []string{
			"1",
			"2",
		}, // Named Field gives flexiblity to not use null
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions)

	anotherDoctor := aDoctor

	fmt.Println(anotherDoctor.actorName)
	anotherDoctor.actorName = "Tom Baker"
	fmt.Println(aDoctor.actorName)
	fmt.Println(anotherDoctor.actorName)

	var firstBird Bird

	firstBird.Animal.Name = "Emu"         // if we want to mimic inheritance
	firstBird.Animal.Origin = "Australia" // if we want to mimic inheritance
	firstBird.CanFly = false
	firstBird.SpeedKPH = 48

	fmt.Println(firstBird)

	var secondBird Bird

	secondBird.Name = "Peacock"
	secondBird.Origin = "India"
	secondBird.CanFly = true
	secondBird.SpeedKPH = 24

	fmt.Println(secondBird)

	thirdBird := Bird{
		Animal:   Animal{Name: "Penguin", Origin: "Alaska"},
		SpeedKPH: 21,
		CanFly:   false,
	}

	fmt.Println(thirdBird)

	tag := reflect.TypeOf(Animal{})
	field, _ := tag.FieldByName("Name")
	fmt.Println(field.Tag)

	// Control Flow
	//If Statement

	number := 50
	guess := 30
	if guess < 1 || !(returnTrue()) || guess > 100 {
		fmt.Println("The guess must be between 1 and 100 !")
	} else if guess > 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Low")
		} else if guess > number {
			fmt.Println("High")
		} else if guess == number {
			fmt.Println("You have got it")
		}
		fmt.Println(guess < number, guess > number, guess == number)
	}
	if validateNumber(0.456789) {
		fmt.Println("They are the same")
	} else {
		fmt.Println("They are the different")
	}
	// Switch Statements
	// break is implicitly implied
	// fallthrough to override break statement
	switchNumber := 10
	switch switchNumber {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4, 5:
		fmt.Println("three, four or five")
	//case switchNumber <= 10:
	//		fmt.Println("less than 20")
	default:
		fmt.Println("other number")
	}

	var interF interface{} = 1.0
	switch interF.(type) {
	case int:
		fmt.Println("It is of type int")
	case float64:
		fmt.Println("It is of type float64")
	case string:
		fmt.Println("It is of type String")
	default:
		fmt.Println("It is of other type")
	}

	// Looping

	//basic loop

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// multiple variables
	// increment operation is not an expression, but a statement

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Printf("i = %d, j=%d \n", i, j)
	}

	// Go doesn't have do while loop
	// Here is the alternative

	loopingVar := 0 // Initialize the variable out gives Scope at the function level and not at loop level

	for loopingVar < 5 {
		fmt.Println(loopingVar)
		loopingVar++
	}
	fmt.Printf("Printing the loopingVar outside the loop : %d\n", loopingVar)

	//do while infinte statements
	fmt.Println("Do While Logic")
	loopingVar = 0
	for {
		if loopingVar == 5 {
			loopingVar++
			continue
		}
		if loopingVar == 10 {
			break
		}
		fmt.Println(loopingVar)
		loopingVar++
	}

	//Nested Looping
	fmt.Println("Nested Loops")
Loop: //Labels
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// Collections

	collectionS := []int{1, 2, 3}
	fmt.Println(collectionS)
	//for range loop
	for k, v := range collectionS {
		fmt.Println(k, v)
	}
	fmt.Println("Printing statePopulations using for ranges")
	for k, v := range statePopulations {
		fmt.Printf("State: %s\tPopulation: %d\n", k, v)
	}

	//Printing only Keys

	for k := range statePopulations {
		fmt.Printf("Keys - State: %s\n", k)
	}

	//Printing only values

	for _, v := range statePopulations {
		fmt.Printf("Values - Population: %d\n", v)
	}

	// Control Flow
	// Defer
	// Panic - Error happens, throw the exception
	// Recover - Recover ( Handle the error)

	//Defer - Last In First Out

	defer fmt.Println("start")
	defer fmt.Println("middle") //Defer
	defer fmt.Println("end")

	// Deferring Variable

	loopingVar = 0
	fmt.Printf("Value before Defering = %d\n", loopingVar)
	defer fmt.Printf("Value in deferring call = %d\n", loopingVar)
	loopingVar = 10
	fmt.Printf("Value after Defering = %d\n", loopingVar)

	fmt.Println("Calling ReadRobots")
	readRobots() //commenting as too much noise.

	// Pointers
	// Functions
	// capital in the first letter means public scope
	fmt.Println("")
	fmt.Println("Functions")
	fmt.Println("")
	sayMessage("Hello Go!!!")
	sayGreeting("Hello", "Stacey")

	sumResult := sumInt(1, 2, 3, 4, 5)
	fmt.Println("The Sum is : ", sumResult)

	divideResult, err := divideFloat(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nThe result of Division is :%f\n", divideResult)

	runtime.GOMAXPROCS(100) // By default we get 4

	wg.Add(1)
	go sayHello()
	wg.Wait()

	/*zeroResult, err := divideFloat(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}*/

	// go Routines
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHelloCounter(i)
		go increment(i)
	}
	wg.Wait()

	//go Routines with Mutex

	for i := 0; i < 10; i++ {
		wg.Add(2)
		mutex.RLock()
		go sayHelloCounterwMutex(i)
		mutex.Lock()
		go incrementwMutex(i)
	}
	wg.Wait()

	//ToDo experiment with more Methods
	// Understand Error Handling
	// Understand Race Condition
	//Interface

	// Channels

	var waitGroup = sync.WaitGroup{}
	channel1 := make(chan int) // only integer types would be sent by this channel.

	// Simple Go routine with channel

	waitGroup.Add(2)
	go func() { // Receiving routine
		i := <-channel1 // Receives the message from the channel
		fmt.Println(i)
		waitGroup.Done()
	}()
	go func() { // Sending routine
		channel1 <- 42 // Sending the message via the channel
		waitGroup.Done()
	}()
	waitGroup.Wait()

	// Complex Go Routine
	var waitGroup2 = sync.WaitGroup{}
	channel2 := make(chan int)

	for j := 0; j < 5; j++ {
		waitGroup2.Add(2)
		go func() { // Receiving routine
			i := <-channel2 // Receives the message from the channel
			fmt.Println(i)
			waitGroup2.Done()
		}()
		go func() { // Sending routine
			channel2 <- 36 // Sending the message via the channel
			waitGroup2.Done()
		}()
	}
	waitGroup2.Wait()

	// Reading and writing can be done in same routine, but it is not advisable.

	//Buffer Channel with controls on reading and writing
	var waitGroup3 = sync.WaitGroup{}
	channel3 := make(chan int)
	waitGroup3.Add(2)
	go func(channel3 <-chan int) { // Receiving routine
		for i := range channel3 {
			fmt.Println(i)
		}
		waitGroup3.Done()
	}(channel3)
	go func(channel3 chan<- int) { // Sending routine
		channel3 <- 22  // Sending the message via the channel
		channel3 <- 152 // Sending the message via the channel
		close(channel3)
		waitGroup3.Done()
	}(channel3)
	waitGroup3.Wait()

}
func sayMessage(msg string) {
	fmt.Println(msg)
}
func sayGreeting(msg, name string) {
	fmt.Println(msg, name)
}

// Input range of integers
// Output int
func sumInt(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func divideFloat(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot provide zero as second value")
	}
	return a / b, nil
}
func returnTrue() bool {
	fmt.Println("Returning True")
	return true
}

func validateNumber(number float64) bool {
	return math.Abs(number/math.Pow(math.Sqrt(number), 2)-1) < 0.001
}

func readRobots() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots)
}
