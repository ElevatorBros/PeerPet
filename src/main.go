package main

import (
	"fmt"
)

func main() {
	fmt.Printf("no errors")
	pet := NewPet("john")

	array := []Pet{*pet}

	_ = WritePetToJson(array)
}
