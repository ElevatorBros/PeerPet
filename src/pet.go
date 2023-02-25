package main

import (
	"encoding/json"
	"fmt"
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
	Constituton  int `json:"constition"`
	Intelligence int `json:"intelligence"`

	Dob time.Time `json:"dob"`
}

func NewPet(name string) *Pet {
	pet := new(Pet)

	pet.Name = name
	pet.Dob = time.Now()

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

func WritePetToJson(pet *Pet, path string) error {
	data, _ := Jsonify(pet)
	//handle error
	err := os.WriteFile(fmt.Sprintf("%s/%s.json", path, pet.Name), data, fs.FileMode(0644))

	return err
}

func Jsonify(pet *Pet) (petJson []byte, e error) {
	return json.Marshal(pet)
}

func UnJsonify(data []byte, pet *Pet) error {
	return json.Unmarshal(data, pet)
}

func ReadPets(folder string) []Pet {
	var pets = []Pet{}
    f, err := os.Open(folder)
    if err != nil {
        panic(err)
    }
    files, err := f.ReadDir(0)
    if err != nil {
        panic(err)
    }
    for _, filename := range files {
        file, err := os.Open(fmt.Sprintf("%s/%s", folder, filename.Name()))
        defer file.Close()
        if err != nil {
            panic(err)
        }
        var data = []byte{}
        pet := new(Pet)
        file.Read(data)
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
