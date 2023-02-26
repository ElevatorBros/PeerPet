package main

import (
	"log"
	"sync"
	"time"

	"github.com/schollz/croc/v9/src/comm"
)

// define 4 hours in seconds
const FOUR_HOURS = 999999999999

// version number
var version_number = 0.0

// GO_Routine sync
var wg sync.WaitGroup

// relay
var relay *Relay

// connection
var connection *comm.Comm

// my pet
var my_pet Pet

// if in combat
var in_combat = true

func main() {
	if relay == nil {
		relay = NewRelay()
	}

	temp_conn, err := relay.JoinRoom("abc123")
	if err != nil {
		panic(err)
	}
	connection = temp_conn

	my_pet = ReadPets()[0]
	my_pet.Print()

	var opponent_pet Pet

	HostCombat(&opponent_pet)
	Combat()
}

func HostCombat(opponent_pet *Pet) {
	UnJsonify(WaitForReceive(FOUR_HOURS), opponent_pet)
	SendPet()
}

func JoinCombat() {
	SendPet()
	data := WaitForReceive(5)
	if data != nil {
		// IMPLEMENT IN GUI TO TELL THAT ROOM IS INVALID
	}
}

func SendPet() {
	data, err := Jsonify(my_pet)
	if err != nil {
		panic(err)
	}

	err = connection.Send(data)
	if err != nil {
		log.Fatal(err)
	}
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

	data := WaitForReceive(FOUR_HOURS)
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

func WaitForReceive(duration float64) []byte {
	start := time.Now()
	var data []byte
	var err error
	for len(data) == 1 || len(data) == 0 {
		data, err = connection.Receive()
		if err != nil {
			log.Fatal(err)
		}
		if time.Since(start).Seconds() < duration {
			return nil
		}
	}
	return data
}
