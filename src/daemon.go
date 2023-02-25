package main

import (
	"fmt"
	"os"
)

func dMain() {
	CreateDataDir()
}

func CreateDataDir() string {
	xdg_data := os.Getenv("XDG_DATA_HOME")
	if xdg_data == "" {
		xdg_data = "~/.local/share"
	}
    os.MkdirAll(fmt.Sprintf("%s/peerpet", xdg_data), os.FileMode(0755))
	return fmt.Sprintf("%s/pets.json", xdg_data)
}
