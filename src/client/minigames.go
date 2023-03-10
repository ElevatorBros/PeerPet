package client

import (
	"log"
	"math/rand"
	"os"
	"strings"

	tc "github.com/gdamore/tcell/v2"
	tv "github.com/rivo/tview"
	"peer.pet/src/common"
)

// typing
func typing(pet *common.Pet) *tv.InputField {
	quote_file, err := os.ReadFile("./rec/quotes.txt")
	if err != nil {
		panic(err)
	}

	quote_string := string(quote_file[:])
	quotes := strings.Split(quote_string, "\n")
	quote := quotes[rand.Intn(len(quotes))]
	var lastMessage string
	i := 0
	log.Print(quote)

	quote = "your mom"
	inputField := tv.NewInputField()
	inputField.SetLabel("Typing!")
	inputField.SetChangedFunc(func(text string) {
		if len(lastMessage) < len(text) {
			i++
		} else {
			i--
		}
		lastMessage = text
	})
	inputField.SetAcceptanceFunc(func(textToCheck string, lastChar rune) bool {
		if i < len(quote) {
			return lastChar == rune(quote[i])
		} else {
			return false
		}
	})
	inputField.SetDoneFunc(func(key tc.Key) {
		switch key {
		case tc.KeyEnter:
			// Change pet Intelligence
			// pet.Intelligence += 2
		case tc.KeyEscape:
			// The tamagatchi will tell the user that they failed
		}
	})

	return inputField
}
