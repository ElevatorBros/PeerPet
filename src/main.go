package main

func main() {
	RunGUI()
	//creates pet named john
	pet := NewPet("john")
	pet1 := NewPet("ronan")
	pet2 := NewPet("dog")
	//makes sure storing directory exists and returns path
	folder_path = CreateDataDir()

	//serializes pet to json
	err := WritePetToJson(pet)
	WritePetToJson(pet1)
	WritePetToJson(pet2)
	if err != nil {
		panic(err)
	}

	//reads stored json files to pet array
	//pets := ReadPets()

	////prints pet array data
	//for _, thing := range pets {
	//	thing.Print()
	//}
}
