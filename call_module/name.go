package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
	 // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    names := []string{
        "Paul",
        "Jessica",
        "Stilgar",
    }

    // Request a greeting message.
    message, err := greetings.Hello("Tim")

    messages, errMap := greetings.Hellos(names)

    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }
     if errMap != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the message to console
    fmt.Println(message)
    //If no error was returned, print the map to the console
    fmt.Println(messages)
}