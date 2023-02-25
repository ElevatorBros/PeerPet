package main

import (
    "log"
    "github.com/gdamore/tcell/v2"
    "github.com/bennicholls/burl-E/reximage"
    "unicode/utf8"
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

func draw_xp_image(offsetX int, offsetY int, image reximage.ImageData) {
    for x := 0; x < image.Width; x++ {
        for y := 0; y < image.Height; y++ {
            cell, _ := image.GetCell(x, y)
            forground_cell_color := tcell.NewRGBColor(int32(cell.R_f), int32(cell.G_f), int32(cell.B_f))
            background_cell_color := tcell.NewRGBColor(int32(cell.R_b), int32(cell.G_b), int32(cell.B_b))
            
            color_style := tcell.StyleDefault.Background(background_cell_color).Foreground(forground_cell_color)
            cell_rune, _ := utf8.DecodeRuneInString(string(cell.Glyph))
            display.SetContent(offsetX + x, offsetY + y, cell_rune, nil, color_style)
        }
    }
}

func draw() {
    image, err := reximage.Import("./rec/pet1.xp")
	if err != nil {
		log.Fatalf("%+v", err)
	}

    draw_xp_image(0, 0, image)
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
