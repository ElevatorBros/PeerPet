package main

import (
	"encoding/json"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"time"
)

type Pet struct {
	Name   string  `json:"name"`
	Hunger float32 `json:"hunger"`
	Thirst float32 `json:"thirst"`
	Energy float32 `json:"energy"`

	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constituton  int `json:"constution"`
	Intelligence int `json:"intelligence"`

	Dob        time.Time `json:"dob"`
	SpritePath []string  `json:"spritepath"`
}

func NewPet(name string) *Pet {
	pet := new(Pet)

	pet.Name = name
	pet.Dob = time.Now()
	pet.SpritePath = []string{"aa"}

	pet.Hunger = 25
	pet.Thirst = 25
	pet.Energy = 100

	RandomizeStats(&pet.Strength, &pet.Dexterity, &pet.Constituton, &pet.Intelligence)

	return pet
}

func RandomizeStats(strength *int, dexterity *int, constition *int, intelligence *int) {
	rand.NewSource(time.Now().UnixNano())

	*strength = rand.Intn(100)
	*dexterity = rand.Intn(100)
	*constition = rand.Intn(100)
	*intelligence = rand.Intn(100)
}

func (pet Pet) AdvanceHunger(value float32) {
	time.Sleep(5 * time.Second)
	pet.Hunger += value
}

func (pet Pet) AdvanceThirst(value float32) {
	time.Sleep(5 * time.Second)
	pet.Hunger += value
}

func (pet Pet) AdvanceEnergy(value float32) {
	time.Sleep(5 * time.Second)
	pet.Energy -= value
}

func Jsonify(pet *Pet) (petJson []byte, e error) {
	return json.Marshal(pet)
}

func UnJsonify(data []byte, pet *Pet) error {
	return json.Unmarshal(data, pet)
}

func WritePetToJson(pet *Pet, path string) error {
	data, err := Jsonify(pet)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path+"/"+pet.Name+".json", data, fs.FileMode(0644))
	return err
}

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
		err = UnJsonify(data, pet)
		if err != nil {
			panic(err)
		}

		pets = append(pets, *pet)
	}
	return pets
}

func (pet Pet) Print() {
	log.Printf("%+v\n", pet)
}
