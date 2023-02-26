package server

import (
	"io/fs"
	"os"

	"peer.pet/src/common"
)

var folder_path = CreateDataDir()

// Writes []bytes to file
func WritePetToJson(pet *common.Pet) error {
	data, err := pet.Jsonify()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(folder_path+"/pet.json", Secret(data), fs.FileMode(0644))
	return err
}

// Reads file to []bytes
func ReadPets() []common.Pet {
	var pets = []common.Pet{}

	folder, err := os.Open(folder_path)
	if err != nil {
		panic(err)
	}

	files, err := folder.ReadDir(0)
	if err != nil {
		panic(err)
	}

	for _, filename := range files {
		data, _ := os.ReadFile(folder.Name() + "/" + filename.Name())

		pet := new(common.Pet)
		err = common.UnJsonify(Secret(data), pet)
		if err != nil {
			panic(err)
		}

		pets = append(pets, *pet)
	}
	return pets
}

// Encrypts & Decrypts
func Secret(data []byte) []byte {
	for i, by := range data {
		data[i] = ^by
	}
	return data
}
