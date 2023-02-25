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

func jsonify(pet *Pet) (petJson []byte, e error) {
	return json.Marshal(pet)
}

func Unjsonify(data []byte, pet *Pet) error {
	return json.Unmarshal(data, pet)
}
