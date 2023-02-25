package main

import "log"

func main() {
	pet := NewPet("john")

	path := CreateDataDir()

	err := WritePetToJson(pet, path)
	if err != nil {
		log.Print(err.Error())
	}

	pets := ReadPets(path)

	for i := range pets {
		log.Printf("%s\n", pets[i].Name)
		pets[i].Print()
	}

}
