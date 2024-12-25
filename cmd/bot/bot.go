// Initializes and runs the Discord bot
package main

import (
    "fmt"
    "log"
    "github.com/faulteh/nap-and-go/pkg/greetings"
    "rsc.io/quote"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("bot: ")
    log.SetFlags(0)

    fmt.Println(quote.Go())

    // A slice of names.
    names := []string{"Ardenne", "Sam", "Liz"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)

    // If an error was returned, print it to the console and exit the program.
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(messages)
}
