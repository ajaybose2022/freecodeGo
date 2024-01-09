package main

import "fmt"

func sayHello() {
	fmt.Printf("Hello from another file \n")
	wg.Done()
}

func sayHelloCounter(counter int) {
	fmt.Printf("Hello from another file with counter: %v\n", counter)
	wg.Done()
}

func increment(counter int) {
	counter++
	wg.Done()
}

func sayHelloCounterwMutex(counter int) {
	//mutex.RLock()
	fmt.Printf("Hello from another file with mutexcounter: %v\n", counter)
	mutex.RUnlock()
	wg.Done()
}

func incrementwMutex(counter int) {
	//mutex.Lock()
	counter++
	mutex.Unlock()
	wg.Done()
}
