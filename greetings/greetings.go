package greetings

import "fmt"

// takes a name parameter whose type is string, and returns a string
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    //the := operator is a shortcut for declaring and initializing a variable in one line
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}