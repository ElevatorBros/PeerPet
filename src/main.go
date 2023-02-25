package main

import (
    "github.com/gdamore/tcell/v2"
)

func main() {
	//fmt.Printf("no errors")

    setup_tcell()
    draw()

    for {
        ev := display.PollEvent()
        switch ev := ev.(type) {
            case *tcell.EventKey:
                if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
                    quit_tcell()
                    return;
                }
            case *tcell.EventResize:
                display.Sync()
        }
    }
}
