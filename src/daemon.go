package main

import (
	"fmt"
	"log"
	"os"
)

func dMain() {
	CreateDataDir()
}

func CreateDataDir() string {
	xdg_data := os.Getenv("XDG_DATA_HOME")
	if xdg_data == "" {
        home := os.Getenv("HOME")
        if home == "" { log.Fatalf("Could not create data dir. Try manually creating ~/.local/share/peerpet") }
		xdg_data = fmt.Sprintf("%s/.local/share", home)
	}
    os.MkdirAll(fmt.Sprintf("%s/peerpet", xdg_data), os.FileMode(0755))
	return fmt.Sprintf("%s/peerpet/pets.json", xdg_data)
}
