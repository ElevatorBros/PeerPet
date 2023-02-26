package main

import (
	"log"
	"os"
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

// my pet (can be used in other parts of program maybe)
var my_pet Pet

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

	pets := ReadPets()
	if len(pets) == 0 {
		// MAKE THIS IN GUI
		log.Fatal("YOU HAVE NO PETS")
	}

	// CHANGE 0 TO USER SELECTION
	my_pet = pets[0]

	// HOST OR NOT
	host := false
	if os.Args[1] == "-S" {
		host = true
	}

	// gets opponent's pet
	var opponent_pet Pet
	if host {
		HostCombat(&opponent_pet)
	} else {
		JoinCombat(&opponent_pet)
	}

	opponent_pet.Print()
	Combat()
}

// joins combat server-side
func HostCombat(opponent_pet *Pet) {
	data := WaitForReceive(FOUR_HOURS)
	UnJsonify(data, opponent_pet)
	SendPet()
}

// joins combat client-side
func JoinCombat(opponent_pet *Pet) {
	SendPet()
	data := WaitForReceive(5)
	if data == nil {
		// IMPLEMENT IN GUI TO TELL THAT ROOM IS INVALID
	}
	UnJsonify(data, opponent_pet)
}

// sends my_pet to opponent
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

// main combat function
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

// wait for attack to be received
func ReceiveAttack() []byte {
	// sends done signal at end of function
	defer wg.Done()

	data := WaitForReceive(FOUR_HOURS)
	//time.Sleep(5 * 1000 * time.Millisecond)
	log.Printf("%s\n", string(data[:]))

	return data
}

// send an attack
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

// wait for data to be received
func WaitForReceive(duration float64) []byte {
	start := time.Now()
	var data []byte
	var err error
	for len(data) == 1 || len(data) == 0 {
		data, err = connection.Receive()
		if err != nil {
			log.Fatal(err)
		}
		if time.Since(start).Seconds() > duration {
			return nil
		}
	}
	return data
}
