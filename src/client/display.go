package client

import (
	// "fmt"
	// "log"
	"fmt"
	// "log"

	"github.com/bennicholls/burl-E/reximage"
	tc "github.com/gdamore/tcell/v2"
	tv "github.com/rivo/tview"
)

var app *tv.Application
var currentButton int

const LEFT_SIZE = 40

func DrawXP(display *tv.Table, offsetX int, offsetY int, image reximage.ImageData) {
	for x := 0; x < image.Width; x++ {
		for y := 0; y < image.Height; y++ {
			img_cell, _ := image.GetCell(x, y)
			fg_col := tc.NewRGBColor(int32(img_cell.R_f), int32(img_cell.G_f), int32(img_cell.B_f))

			color_style := tc.StyleDefault.Foreground(fg_col)

			cell := tv.NewTableCell(fmt.Sprintf("%c", img_cell.Glyph))
			cell.SetStyle(color_style)

			display.SetCell(offsetY+y, offsetX-1+x, cell)
		}
	}
}

func TabbableSupport(w *tv.Flex) *tv.Flex {
	currentButton = 0
	w.SetInputCapture(func(e *tc.EventKey) *tc.EventKey {
		if w.HasFocus() {
			// Movement
			switch e.Key() {
			case tc.KeyTab:
				currentButton += 1
				currentButton %= w.GetItemCount()
				app.SetFocus(w.GetItem(currentButton))
			}
		}
		return e
	})
	return w
}

func CombatSetup(pages *tv.Pages) {
	combat_setup := tv.NewFlex()
	combat_host_or_client := tv.NewFlex()
	combat_host_flex := tv.NewFlex()
	combat_client_form := tv.NewInputField()

	TabbableSupport(combat_host_or_client)

	host_button := tv.NewButton("host")
	client_button := tv.NewButton("client")
	combat_host_or_client.AddItem(host_button, 0, 1, false)
	combat_host_or_client.AddItem(client_button, 0, 1, false)

	pages.AddPage("combat_host_or_client", combat_host_or_client, true, false)
	pages.AddPage("combat_setup", combat_setup, true, false)

	host_button.SetSelectedFunc(func() {
        key := GetHostKey()
        view := tv.NewTextView().SetText(key).SetTextAlign(tv.AlignCenter)
        combat_host_flex.AddItem(view, 0, 1, true)
		combat_setup.AddItem(combat_host_flex, 0, 1, true)
        pages.SwitchToPage("combat_setup")
        go func() {
            EnterCombat(true)
            pages.SwitchToPage("combat_actual")
        }()
	})
	client_button.SetSelectedFunc(func() {
        combat_client_form.SetLabel("Enter host's key here:").SetDoneFunc(func(key tc.Key) {
            SetClientKey(combat_client_form.GetText())
            EnterCombat(false)
            pages.SwitchToPage("combat_actual")
        })
		combat_setup.AddItem(combat_client_form, 0, 1, true)
        pages.SwitchToPage("combat_setup")
	})
}

func CombatActual(flex *tv.Flex) {
    txt := tv.NewTextArea()
    input := tv.NewInputField()
    flex.AddItem(txt, 0, 1, false)
    flex.AddItem(input, 0, 1, true)
}

func RunGUI() {
	app = tv.NewApplication()
	main := tv.NewFlex()
	combat_actual := tv.NewFlex()
	pages := tv.NewPages()
	TabbableSupport(main)

	roullette_button := tv.NewButton("Roullette")
	roullette_button.Focus(func(p tv.Primitive) {
		roullette_button.SetBackgroundColor(tc.ColorBlue)
	})

	combat_button := tv.NewButton("Combat")
	combat_button.SetSelectedFunc(func() {
		pages.SwitchToPage("combat_host_or_client")
	})

	main.AddItem(combat_button, 0, 1, true)
	main.AddItem(roullette_button, 0, 1, false)

    CombatSetup(pages)
    CombatActual(combat_actual)
	pages.AddPage("main", main, true, true)
	pages.AddPage("combat_actual", combat_actual, true, false)

	app.SetRoot(pages, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
