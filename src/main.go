package main

var time_to_quit chan struct{}

func main() {
	//creates pet named john
	pet := NewPet("john")
	//makes sure storing directory exists and returns path
	path := CreateDataDir()

	//serializes pet to json
	err := WritePetToJson(pet, path)
	if err != nil {
		panic(err)
	}

	//reads stored json files to pet array
	pets := ReadPets(path)

	//prints pet array data
	for _, thing := range pets {
		thing.Print()
	}
}
