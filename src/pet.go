package main

import (
	"time"
)

type Pet struct {
	name   string
	hunger float32
	thirst float32
	energy float32

	strength     int
	dexterity    int
	constituton  int
	intelligence int

	dob time.Time
    xp_path string
}

func (p Pet) age() {

}
