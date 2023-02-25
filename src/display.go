package main

import (
    "log"
    "github.com/gdamore/tcell/v2"
)


var display tcell.Screen
var def_style tcell.Style

func hello() {
    display.SetContent(0, 0, 'H', nil, def_style)
    display.SetContent(1, 0, 'i', nil, def_style)
    display.SetContent(2, 0, '!', nil, def_style)
}

func quit_tcell() {
    maybePanic := recover()
    display.Fini()
    if maybePanic != nil {
        panic(maybePanic)
    }
}

func draw() {
    display.Show()
}

func setup_tcell() {
    def_style = tcell.StyleDefault
    var err error
    display, err = tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := display.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

    display.SetStyle(def_style)
    display.Clear()
}
