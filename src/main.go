package main

import (
	//"fmt"
    "time"
)

func main() {
	//fmt.Printf("no errors")

    setup_tcell()
    hello()
    draw()
    time.Sleep(1000 * time.Millisecond)
    quit_tcell()
}
