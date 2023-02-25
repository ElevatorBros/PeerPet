package main

import (
	"encoding/json"
	"math/rand"
	"time"
)

type Pet struct {
	Name   string
	Hunger float32
	Thirst float32
	Energy float32

	Strength     int
	Dexterity    int
	Constituton  int
	Intelligence int

	Dob time.Time
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

	//power := rand.Int(400) + 300

}

func Jsonify(pet []Pet) (petJson []byte, e error) {
	return json.Marshal(pet)
}

func Unjsonify(data []byte, pet *Pet) error {
	return json.Unmarshal(data, pet)
}
