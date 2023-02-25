package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

var in_combat = true
var wg sync.WaitGroup

func main() {
	http.Head("a")
	Combat()
}
func Combat() {
	wg.Add(2)
	if in_combat {
		go ReceiveAttack()
		go SendAttack()
	}

	wg.Wait()
	log.Print("finished")
}

func ReceiveAttack() {
	defer wg.Done()
	time.Sleep(5 * 1000 * time.Millisecond)
	log.Print("received")
}

func SendAttack() {
	defer wg.Done()
	time.Sleep(2 * 1000 * time.Millisecond)
	log.Print("sent")

}
