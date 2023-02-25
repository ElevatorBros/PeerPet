package main

import "log"

func main() {
	pet := NewPet("john")

	path := CreateDataDir()

	_ = WritePetToJson(pet, path)

	pets := ReadPets(path)

	for i := range pets {
		log.Printf("%s\n", pets[i].Name)
		pets[i].Print()
	}

}
