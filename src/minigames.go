<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> bc8e1e5 (typing game!)
package main
<<<<<<< HEAD
<<<<<<< HEAD

=======
package main 
/*
>>>>>>> e5565ce (combat buttons)
=======
/*
>>>>>>> 32edd7b (For Ronan)
=======

>>>>>>> aa205e1 (love you josh)
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
<<<<<<< HEAD
<<<<<<< HEAD
func typing(pet *Pet) *tv.InputField {
=======
func typing(w *tv.Flex, pet *Pet) {
>>>>>>> bc8e1e5 (typing game!)
=======
func typing(w *tv.Flex, pet *Pet) *tv.InputField {
>>>>>>> 915ade6 (ronna can edit his code now)
=======
func typing(pet *Pet) *tv.InputField {
>>>>>>> c82890a (something)
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
<<<<<<< HEAD
=======
>>>>>>> 09915fe (typing game error checking)
        if i < len(quote) {
            return lastChar == rune(quote[i])
        } else {
            return false
        }
<<<<<<< HEAD
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
=======
>>>>>>> 09915fe (typing game error checking)
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

<<<<<<< HEAD
    w.AddItem(inputField, 0, 1, true)
<<<<<<< HEAD
>>>>>>> bc8e1e5 (typing game!)
}
<<<<<<< HEAD
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
=======

>>>>>>> aa205e1 (love you josh)
=======
=======
>>>>>>> c82890a (something)
    return inputField
}
>>>>>>> 915ade6 (ronna can edit his code now)
