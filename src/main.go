package main

import ()

func main() {
	pet := NewPet("john")

	path := CreateDataDir()

	array := []Pet{*pet}

	_ = WritePetToJson(array, path)

	pets := ReadPets(path)

	for i := range pets {
		pets[i].Print()
	}

}
