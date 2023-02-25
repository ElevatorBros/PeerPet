package main

import (
	// "fmt"
	// "log"
	// "unicode/utf8"
	"github.com/bennicholls/burl-E/reximage"
	tc "github.com/gdamore/tcell/v2"
	tv "github.com/rivo/tview"
)


var display tv.Application
const LEFT_SIZE = 40 
// var style tv.Styles

func draw_xp_image(offsetX int, offsetY int, image reximage.ImageData) {
    // for x := 0; x < image.Width; x++ {
    //     for y := 0; y < image.Height; y++ {
    //         cell, _ := image.Getc(x, y)
    //         forground_cell_color := tc.NewRGBColor(int32(cell.R_f), int32(cell.G_f), int32(cell.B_f))
    //         background_cell_color := tc.NewRGBColor(int32(cell.R_b), int32(cell.G_b), int32(cell.B_b))
    //         
    //         color_style := tc.StyleDefault.Background(background_cell_color).Foreground(forground_cell_color)
    //         cell_rune, _ := utf8.DecodeRuneInString(string(cell.Glyph))
    //         display.SetContent(offsetX + x, offsetY + y, cell_rune, nil, color_style)
    //     }
    // }
}

func draw_text(x int, y int, text string) {
    // for i := 0; i < utf8.RuneCountInString(text); i++ {
    //     display.SetContent(x + i, y, rune(text[i]), nil, def_style)
    // }
}

func draw(page int, pet Pet) {
    // numbers
    // draw_text(5, 1, "1    2    3    4    5")
    // display.SetContent(5*page + 4, 1, '[', nil, def_style)
    // display.SetContent(5*page + 6, 1, ']', nil, def_style)
    //
    // switch page {
    //     case 0: 
    //         // Pet Image
    //         image, err := reximage.Import("./rec/pet1.xp")
    //         if err != nil {
    //             log.Fatalf("%+v", err)
    //         }
    //         draw_xp_image(0, 3, image)
    //
    //         // Stats
    //         draw_text(20, 4, "     Name : " + pet.Name)
    //         draw_text(20, 6, "Happiness : " + "Fix Orestest")
    //         draw_text(20, 7, "   Hunger : " + fmt.Sprintf("%v", pet.Hunger))
    //         draw_text(20, 8, "   Energy : " + fmt.Sprintf("%v", pet.Energy))
    //         draw_text(20, 9, "   Thirst : " + fmt.Sprintf("%v", pet.Thirst))
    // }
    //
    // display.Show()
}

func Init(input chan tc.Key) {
 //    def_style = tc.StyleDefault
 //    var err error
 //    display, err = tc.NewScreen()
	// if err != nil {
	// 	log.Fatalf("%+v", err)
	// }
	// if err := display.Init(); err != nil {
	// 	log.Fatalf("%+v", err)
	// }
	//
 //    display.SetStyle(def_style)
 //    display.Clear()
	//
 //    go func() {
 //        for {
 //            // Poll event
 //            ev := display.PollEvent()
	//
 //            // Process event
 //            switch ev := ev.(type) {
 //            case *tc.EventResize:
 //                display.Sync()
 //            case *tc.EventKey:
 //                if ev.Key() == tc.KeyEscape || ev.Rune() == 'q' {
 //                    close(time_to_quit)
 //                    return
 //                }
 //                input <- ev.Key()
 //            }
 //        }
 //    }()
}

func RunGUI() {
    pet_box := tv.NewBox().SetBorder(true).SetTitle("Pet")
    message_box := tv.NewBox().SetBorder(true).SetTitle("Message")
    left_flex := tv.NewFlex().SetDirection(tv.FlexRow).
        AddItem(pet_box, 0, 2, false).
        AddItem(message_box, 0, 1, false)

    flex := tv.NewFlex().
        AddItem(left_flex, 0, 1, false)

	if err := tv.NewApplication().SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
