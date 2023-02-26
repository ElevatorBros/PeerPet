package main

import (
	"bufio"
	"log"
	"os"
	"sync"
	"time"

	"github.com/schollz/croc/v9/src/comm"
	"github.com/schollz/croc/v9/src/utils"
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

// TEMPORARY TERMINAL READER
var reader = bufio.NewReader(os.Stdin)

func main() {
	host := false
	if len(os.Args) == 2 && os.Args[1] == "-S" {
		host = true
	}

	if relay == nil {
		relay = NewRelay()
	}

	var key string
	if host {
		key = utils.GetRandomName()
		log.Printf("ROOM KEY: %s\n", key)
	} else {
		log.Print("INPUT ROOM KEY: }")
		text, _ := reader.ReadString('\n')
		key = text
		key = key[:len(key)-1]
	}

	temp_conn, err := relay.JoinRoom(key)
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
	// gets opponent's pet
	var opponent_pet Pet
	if host {
		HostCombat(&opponent_pet)
	} else {
		JoinCombat(&opponent_pet)
	}
	log.Print(key)

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
		go SendAttack()

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
func SendAttack() {
	// sends done signal at end of function
	defer wg.Done()

	text, _ := reader.ReadString('\n')
	data := []byte(text)

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
