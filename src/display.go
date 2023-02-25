package main

import (
    "log"
    "github.com/gdamore/tcell/v2"
    "github.com/bennicholls/burl-E/reximage"
)


var display tcell.Screen
var def_style tcell.Style

func quit_tcell() {
    maybePanic := recover()
    display.Fini()
    if maybePanic != nil {
        panic(maybePanic)
    }
}

func draw_xp_image(image reximage.ImageData) {
    for x := 0; x < image.Width; x++ {
        for y := 0; y < image.Height; y++ {
        }
    }
}

func draw() {
    image, err := reximage.Import("./rec/test.xp")
	if err != nil {
		log.Fatalf("%+v", err)
	}

    draw_xp_image(image)
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
