package main 

import (
    "os"
    "strings"
    tc "github.com/gdamore/tcell/v2"
    tv "github.com/rivo/tview"
)

var quote_index = -1

func get_input();

// typing
func typing(window tv.Box, pet *Pet) {
    quote_file, _ := os.ReadFile("./rec/quotes.txt")
    quote_string := string(quote_file[:])
    quotes := strings.Split(quote_string, "\n")
    if quote_index == -1 {
        quote_index = 1
    }

    quote := quotes[quote_index]
    var input rune[]
    for 1 {
        missing := false
        for int c = 0; c < len(quote); c++ {
            color = white
            if c < len(input) {
                if quote[c] != input[c] {
                    color = red
                    missing = true
                } else {
                    color = blue
                }
            }
            display(quote[c])
        }

        if !missing && len(input) == len(quote) {
            break
        }
       
        in := get_input()
        if in == backspace {
            input = input[:len(input)-1]
        } else if ascii {
            input = append(input, in)
        }

    }
}
