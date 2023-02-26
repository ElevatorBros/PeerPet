package main

import (
	// "fmt"
	// "log"
	"fmt"

	"github.com/bennicholls/burl-E/reximage"
	tc "github.com/gdamore/tcell/v2"
	tv "github.com/rivo/tview"
)


var app tv.Application
const LEFT_SIZE = 40 

func DrawXP(display *tv.Table, offsetX int, offsetY int, image reximage.ImageData) {
    for x := 0; x < image.Width; x++ {
        for y := 0; y < image.Height; y++ {
            img_cell, _ := image.GetCell(x, y)
            fg_col := tc.NewRGBColor(int32(img_cell.R_f), int32(img_cell.G_f), int32(img_cell.B_f))
            
            color_style := tc.StyleDefault.Foreground(fg_col)

            cell := tv.NewTableCell(fmt.Sprintf("%c", img_cell.Glyph))
            cell.SetStyle(color_style)

            display.SetCell(offsetY + y, offsetX - 1 + x, cell)
        }
    }
}

func drawPet(page int, pet Pet) {
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

func MonitorInput(e *tc.EventKey) *tc.EventKey {
    // Put your funny code here josh

    return e;
}


func RunGUI() {
    app := tv.NewApplication()

    pet_table := tv.NewTable()
    pet_table.SetBorder(true).SetTitle("Pet")
    
    message_box := tv.NewBox()
    message_box.SetBorder(true).SetTitle("Message")
    
    stats_box := tv.NewTable()
    stats_box.SetBorder(true).SetTitle("Stats")

    left_flex := tv.NewFlex()
    left_flex.SetDirection(tv.FlexRow).
        AddItem(pet_table, 0, 5, false).
        AddItem(message_box, 0, 1, false).
        AddItem(stats_box, 0, 4, false)

    /*
    game_grid := tv.NewGrid()
    game_grid.SetColumns(8)
    game_grid.SetBorder(true).SetTitle("Games")

    game_grid.AddItem(tv.NewTextView().
			SetTextAlign(tv.AlignCenter).
			SetText("Typing"), 0, 0, 1, 3, 0, 0, false)


    game_grid.AddItem(tv.NewTextView().
			SetTextAlign(tv.AlignCenter).
			SetText("Roullette"), 0, 4, 1, 3, 0, 0, false).
            SetBorderColor(tc.ColorRed)

    context := 0
    game_grid.SetInputCapture(func(e *tc.EventKey) *tc.EventKey {
		// Movement
		switch e.Key() {
		case tc.KeyTab:
            context += 1
		}

        game_grid.SetTitleColor(tc.ColorRed)

		return e
	})
    */
    game_flex := tv.NewFlex()
    game_flex.SetBorder(true).SetTitle("Games")


    typing_button := tv.NewButton("Typing")
    typing_button.SetInputCapture(func(e *tc.EventKey) *tc.EventKey {
        if game_flex.HasFocus() {
            // Movement
            switch e.Key() {
            case tc.KeyEnter:
                tmp := NewPet("hi")
                field := typing(tmp)
                game_flex.AddItem(field, 0, 2, true)
            }
        }


		return e
	})

    roullette_button := tv.NewButton("Roullette")
    roullette_button.Focus(func(p tv.Primitive) {
        roullette_button.SetBackgroundColor(tc.ColorBlue)
    })

    game_flex.AddItem(typing_button, 0, 1, false)
    game_flex.AddItem(roullette_button, 0, 1, false)


    context := 0
    game_flex.SetInputCapture(func(e *tc.EventKey) *tc.EventKey {
        if game_flex.HasFocus() {
            // Movement
            switch e.Key() {
            case tc.KeyTab:
                context += 1
                context %= game_flex.GetItemCount()
                app.SetFocus(game_flex.GetItem(context))
            }
        }


		return e
	})
    
    right_flex := tv.NewFlex()
    right_flex.SetDirection(tv.FlexRow).
        AddItem(game_flex, 0, 1, false)

    image, _ := reximage.Import("./rec/pet1.xp")
    offsetY, offsetX, _, _ := pet_table.GetInnerRect()

    DrawXP(pet_table, offsetX, offsetY, image)

    flex := tv.NewFlex().
        AddItem(left_flex, 0, 2, false).
        AddItem(right_flex, 0, 3, false)

    flex.SetBackgroundColor(tc.ColorDefault)



	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
