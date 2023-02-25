package main

import (
	"log"
	"os"
)

func dMain() {
    xdg_data := os.Getenv("XDG_DATA_HOME")
    if xdg_data == "" { 
      xdg_data = "~/.local/share" 
    }
    folder, err := os.Stat(xdg_data)
    if err != nil { 
      log.Fatalf("Error opening folder: %v", folder) 
    }
    if !folder.IsDir() {
      log.Fatalf("Folder is not directory")
    }
}
