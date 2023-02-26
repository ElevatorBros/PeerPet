<<<<<<< HEAD:src/client/combat.go
package client

=======
package main
/*
>>>>>>> 32edd7b (For Ronan):src/combat.go
import (
	"bufio"
	"log"
	"os"
	"sync"
	"time"

	"github.com/schollz/croc/v9/src/comm"
	"github.com/schollz/croc/v9/src/utils"
	"peer.pet/src/common"
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
var my_pet *common.Pet

// if in combat
var in_combat = true

// TEMPORARY TERMINAL READER
var reader = bufio.NewReader(os.Stdin)

func EnterCombat(host bool) {
	_, key := InitializeCombat()

	var err error
	// select my_pet
	my_pet, err = ReadPet()
	if err != nil {
		panic(err)
	}
	// gets opponent's pet based on host or not
	var opponent_pet common.Pet
	if host {
		HostCombat(&opponent_pet)
	} else {
		JoinCombat(&opponent_pet)
	}
	log.Print(key)

	log.Print("You are against this pet: ")
	opponent_pet.Print()

	Combat()
}

func InitializeCombat() (bool, string) {
	// set host based on command args
	host := false
	if len(os.Args) == 2 && os.Args[1] == "-S" {
		host = true
	}

	// initialize relay if not ready
	if relay == nil {
		relay = NewRelay()
	}

	// dispay or request room key
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

	// join room
	temp_conn, err := relay.JoinRoom(key)
	if err != nil {
		panic(err)
	}
	connection = temp_conn

	return host, key
}

// joins combat server-side
func HostCombat(opponent_pet *common.Pet) {
	data := WaitForReceive(FOUR_HOURS)
	common.UnJsonify(data, opponent_pet)
	SendPet()
}

// joins combat client-side
func JoinCombat(opponent_pet *common.Pet) {
	SendPet()
	data := WaitForReceive(5)
	if data == nil {
		// IMPLEMENT IN GUI TO TELL THAT ROOM IS INVALID
	}
	common.UnJsonify(data, opponent_pet)
}

// sends my_pet to opponent
func SendPet() {
	data, err := my_pet.Jsonify()
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

		go ReceiveAttack(&data)
		go SendAttack()

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
	//time.Sleep(5 * 1000 * time.Millisecond)
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
	log.Printf("DAMAGE SENT: %s\n", string(data[:]))
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
*/