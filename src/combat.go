package main

import (
	"log"
	"sync"
	"time"
)

// if in combat
var in_combat = true

// GO_Routine sync
var wg sync.WaitGroup

func main() {
	Combat()
}
func Combat() {
	//number of goroutine to wait to finish is 2
	wg.Add(2)

	if in_combat {
		go ReceiveAttack()
		go SendAttack()
	}

	// waits for goroutines to send done signal
	wg.Wait()
	log.Print("finished")
}

func ReceiveAttack() {
	// sends done signal at end of function
	defer wg.Done()

	time.Sleep(5 * 1000 * time.Millisecond)
	log.Print("received")
}

func SendAttack() {
	// sends done signal at end of function
	defer wg.Done()

	time.Sleep(2 * 1000 * time.Millisecond)
	log.Print("sent")

}
