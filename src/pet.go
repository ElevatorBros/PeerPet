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

	power := RandomizeStats()

	return pet
}

func RandomizeStats(p) {
	val := 0
	return val
}

func Jsonify(pet *Pet) (petJson []byte, e error) {
	return json.Marshal(pet)
}

func Unjsonify(data []byte, pet *Pet) error {
	return json.Unmarshal(data, pet)
}
