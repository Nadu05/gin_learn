package main

import (
	"fmt"
	"time"
)

func greeting(name string) {
	fmt.Printf("Hello, %v!", name)
}

func slowGreeting(name string, doneCh chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Hello, %v!", name)
	doneCh <- true
	close(doneCh)

}
func main() {

	dones := make([]chan bool, 3)
	//doneCh := make(chan bool) //channel -> this like as communication tunnel
	//dones[0] = make(chan bool)
	go slowGreeting("Bro", dones[0])

	//dones[1] = make(chan bool)
	go slowGreeting("Germany", dones[1])
	//dones[2] = make(chan bool)
	go slowGreeting("Brazil", dones[2])

	//for index, done := range dones {
	//	<-done
	//	fmt.Printf("%d: %t\n", index, done)
	//}

	for range dones {

	}

}
