package main

import (
	"errors"
	"os"
)

func dMain() {
	CreateDataDir()
}

func CreateDataDir() {
	xdg_data := os.Getenv("XDG_DATA_HOME")
	if xdg_data == "" {
		xdg_data = "~/.local/share"
	}
	folder, err := os.Stat(xdg_data)
	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir(folder.Name(), os.ModeDir)
	}
}
