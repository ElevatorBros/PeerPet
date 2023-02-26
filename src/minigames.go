<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> bc8e1e5 (typing game!)
package main

=======
package main 
/*
>>>>>>> e5565ce (combat buttons)
import (
	"log"
	"math/rand"
	"os"
	"strings"

	tc "github.com/gdamore/tcell/v2"
	tv "github.com/rivo/tview"
)

// typing
<<<<<<< HEAD
func typing(pet *Pet) *tv.InputField {
=======
func typing(w *tv.Flex, pet *Pet) {
>>>>>>> bc8e1e5 (typing game!)
    quote_file, err := os.ReadFile("./rec/quotes.txt")
    if err != nil { panic(err) }

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
        if len(lastMessage) < len(text) { i++ } else { i-- }
        lastMessage = text
    })
    inputField.SetAcceptanceFunc(func(textToCheck string, lastChar rune) bool {
<<<<<<< HEAD
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
=======
        return lastChar == rune(quote[i])
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

    w.AddItem(inputField, 0, 1, true)
>>>>>>> bc8e1e5 (typing game!)
}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======

>>>>>>> aa205e1 (love you josh)
=======
>>>>>>> 915ade6 (ronna can edit his code now)
=======
*/
>>>>>>> e5565ce (combat buttons)
