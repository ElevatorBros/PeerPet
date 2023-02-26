package main

import (
	"log"
	"sync"

	"github.com/schollz/croc/v9/src/comm"
)

// version number
var version_number = 0.0

// GO_Routine sync
var wg sync.WaitGroup

// relay
var relay *Relay

// connection
var connection *comm.Comm

// if in combat
var in_combat = true

func main1() {
	if relay == nil {
		relay = NewRelay()
	}

	temp_conn, err := relay.JoinRoom("abc123")
	if err != nil {
		panic(err)
	}
	connection = temp_conn

	my_pet := ReadPets()[0]
	my_pet.Print()

	SendPets()
	Combat()
}

func HostCombat() {

	//wait for a pet reception
	//send own pet

}

func JoinCombat() {
	//once joined, send pet
	//if I don't receive pet in k seconds -> invalid room

}

func SendPets() {

}

func Combat() {
	defer connection.Close()

	for in_combat {
		//number of goroutine to wait to finish is 2
		wg.Add(2)

		go ReceiveAttack()
		go SendAttack([]byte("A message"))

		// waits for goroutines to send done signal
		wg.Wait()
		log.Print("finished")
	}
}

func ReceiveAttack() []byte {
	// sends done signal at end of function
	defer wg.Done()

	data := WaitForReceive()
	//time.Sleep(5 * 1000 * time.Millisecond)
	log.Printf("%s\n", string(data[:]))

	return data
}

func SendAttack(data []byte) {
	// sends done signal at end of function
	defer wg.Done()

	err := connection.Send(data)
	if err != nil {
		log.Fatal(err)
	}

	//time.Sleep(2 * 1000 * time.Millisecond)
	log.Printf("%s\n", string(data[:]))
}

func WaitForReceive() []byte {
	var data []byte
	var err error
	for len(data) == 1 || len(data) == 0 {
		data, err = connection.Receive()
		if err != nil {
			log.Fatal(err)
		}
	}
	return data
}
