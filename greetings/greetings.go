package greetings

import (
	"fmt"
	"errors"
) 

// takes a name parameter whose type is string, and returns a string
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("name value empty")
	}
    // Return a greeting that embeds the name in a message.
    //the := operator is a shortcut for declaring and initializing a variable in one line
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}