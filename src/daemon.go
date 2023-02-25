package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func dMain() {
    createDataDir()
}

func createDataDir() {
    xdg_data := os.Getenv("XDG_DATA_HOME")
    if xdg_data == "" { 
      xdg_data = "~/.local/share" 
    }
    folder, err := os.Stat(xdg_data)
    if errors.Is(err, os.ErrNotExist) {
        if err := os.Mkdir(folder.Name(), os.ModePerm); err != nil {
            log.Fatalln("~/.local/share does not exist! Manually create it")
        }
    }
    if !folder.IsDir() {
      log.Fatalf("Folder is not directory")
    }

    var pet Pet
    readPets(fmt.Sprintf("%s/pets.json", xdg_data), &pet)
}

func readPets(filename string, pet *Pet) {
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
    if err != nil  {
        log.Fatalf("Error reading json data file")
    }
}
