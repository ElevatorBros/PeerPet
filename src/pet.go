package main

import (
	"log"
	"math/rand"
	"time"
)

// Pet Structure
type Pet struct {
	Name   string  `json:"name"`
	Hunger float32 `json:"hunger"`
	Thirst float32 `json:"thirst"`
	Energy float32 `json:"energy"`

	Strength     int `json:"strength"`
	Defense      int `json:"defense"`
	Constituton  int `json:"constution"`
	Intelligence int `json:"intelligence"`

	Dob        time.Time `json:"dob"`
	SpritePath []string  `json:"spritepath"`
}

// Pet Constructor
func NewPet(name string) *Pet {
	pet := new(Pet)

	pet.Name = name
	pet.Dob = time.Now()
	pet.SpritePath = []string{"aa"}

	pet.Hunger = 25
	pet.Thirst = 25
	pet.Energy = 100

	RandomizeStats(&pet.Strength, &pet.Defense, &pet.Constituton, &pet.Intelligence)

	return pet
}

// Pet Stat Randomizer
func RandomizeStats(strength *int, dexterity *int, constition *int, intelligence *int) {
	rand.NewSource(time.Now().UnixNano())

	*strength = rand.Intn(100)
	*dexterity = rand.Intn(100)
	*constition = rand.Intn(100)
	*intelligence = rand.Intn(100)
}

// Progresses Hunger
func (pet Pet) ProgressHunger(value float32) {
	time.Sleep(5 * time.Second)
	pet.Hunger += value
}

// Progresses Thirst
func (pet Pet) ProgressThirst(value float32) {
	time.Sleep(5 * time.Second)
	pet.Thirst += value
}

// Progresses Energy
func (pet Pet) ProgressEnergy(value float32) {
	time.Sleep(5 * time.Second)
	pet.Energy -= value
}

// Pet ToString
func (pet Pet) Print() {
	log.Printf("%+v\n", pet)
}
