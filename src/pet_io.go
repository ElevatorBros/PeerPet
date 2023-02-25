package main

import (
	"encoding/json"
	"io/fs"
	"os"
)

// Converts Pet to []bytes
func Jsonify(pet *Pet) (petJson []byte, e error) {
	return json.Marshal(pet)
}

// Converts []bytes to Pet
func UnJsonify(data []byte, pet *Pet) error {
	return json.Unmarshal(data, pet)
}

// Writes []bytes to file
func WritePetToJson(pet *Pet, path string) error {
	data, err := Jsonify(pet)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path+"/"+pet.Name+".json", Secret(data), fs.FileMode(0644))
	return err
}

// Reads file to []bytes
func ReadPets(folder_path string) []Pet {
	var pets = []Pet{}

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

		pet := new(Pet)
		err = UnJsonify(Secret(data), pet)
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
