package main

import (
	"log"
)

var time_to_quit chan struct{}

func main() {
	pet := NewPet("john")
	path := CreateDataDir()

	err := WritePetToJson(pet, path)
	if err != nil {
		log.Print(err.Error())
	}

	pets := ReadPets(path)

	for _, thing := range pets {
		thing.Print()
	}

	/*
		setup_tcell()
		time_to_quit = make(chan struct{})

		for {
			select {
			case _, ok := <-time_to_quit:
				if ok == false {
					return
				}
			case <-time.After(time.Millisecond * 1):
				break
			}

			// draw
			draw(0, *pet)

			get_input()
		}
	*/
}
