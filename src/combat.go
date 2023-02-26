package main

import (
	"log"
	"sync"
	"time"

	"github.com/schollz/croc/v9/src/comm"
	"github.com/schollz/croc/v9/src/utils"
)

// define 4 hours in seconds
const FOUR_HOURS = 14400

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

// room key
var key string

func EnterCombat(host bool) {

	//SET OR GET KEY BEFORE CALLING THIS FUNCTION

	InitializeCombat()

	InitializePets(host)

	Combat()
}

func GetHostKey() string {
	key = utils.GetRandomName()
	return key
}

func SetClientKey(val string) {
	key = val
}

func InitializeCombat() {
	// initialize relay if not ready
	if relay == nil {
		relay = NewRelay()
	}

	// join room
	temp_conn, err := relay.JoinRoom(key)
	if err != nil {
		panic(err)
	}
	connection = temp_conn
}

func InitializePets(host bool) {
	// select my_pet
	// DO SERVER GET PET INSTEAD OF THIS
	pets := ReadPets()
	if len(pets) == 0 {
		// MAKE THIS IN GUI
		log.Fatal("YOU HAVE NO PETS")
	}

	// CHANGE 0 TO USER SELECTION
	my_pet = pets[0]

	// THIS IS BLOCKING
	// gets opponent's pet
	var opponent_pet Pet
	if host {
		HostCombat(&opponent_pet)
	} else {
		JoinCombat(&opponent_pet)
	}
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

	var data []byte
	for in_combat {
		//number of goroutine to wait to finish is 2
		wg.Add(2)

		// PASS IN THE CONTENT OF THE RECEIVE LABEL TO THIS
		go ReceiveAttack(&data)

		// THE ATTACK BUTTON CALLS THIS
		go SendAttack("Your mom")

		// waits for goroutines to send done signal
		wg.Wait()

		log.Printf("DAMAGE RECEIVED: %s\n", string(data[:]))
		CalculateRound()
	}
}

// calculates numbers in round
func CalculateRound() {
	log.Print("finished")
}

// wait for attack to be received
func ReceiveAttack(data *[]byte) {
	// sends done signal at end of function
	defer wg.Done()

	*data = WaitForReceive(FOUR_HOURS)
}

// send an attack
func SendAttack(attack string) {
	// sends done signal at end of function
	defer wg.Done()

	data := []byte(attack)

	err := connection.Send(data)
	if err != nil {
		log.Fatal(err)
	}
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
