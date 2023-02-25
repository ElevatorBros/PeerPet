package main

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
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

func Jsonify(pet []Pet) (petJson []byte, e error) {
	return json.Marshal(pet)
}

func WritePetToJson(pet []Pet) error {
	data, _ := Jsonify(pet)
	path := CreateDataDir()
	err := os.WriteFile(path, data, fs.FileMode(0644))

	return err
}

func Unjsonify(data []byte, pet *Pet) error {
	return json.Unmarshal(data, pet)
}

func ReadPets(filename string, pet *Pet) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatalf("Error reading pets.json file")
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading pets.json file")
	}
	err = Unjsonify(b, pet)
	if err != nil {
		log.Fatalf("Error reading json data file")
	}
}

func (pet Pet) Print() {
	log.Printf("%+v\n", pet)
}
