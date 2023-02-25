package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
    createDataDir()
}

func createDataDir() {
    xdg_data := os.Getenv("XDG_DATA_HOME")
    if xdg_data == "" { 
      xdg_data = "~/.local/share" 
    }
    xdg_data = fmt.Sprintf("%s/peerpet", xdg_data)
    folder, err := os.Stat(xdg_data)
    if errors.Is(err, os.ErrNotExist) {
        os.Mkdir(folder.Name(), os.ModeDir)
    }
}
