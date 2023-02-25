package main

import (
	"encoding/json"
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

	return pet
}

func Randomize() int {

	return val
}

func Jsonify(pet *Pet) (petJson []byte, e error) {
	return json.Marshal(pet)
}
