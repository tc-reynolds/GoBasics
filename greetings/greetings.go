package greetings

import (
	"fmt"
	"errors"

	"math/rand"
    "time"
) 

// takes a name parameter whose type is string, and returns a string
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("name value empty")
	}
    // Return a greeting that embeds the name in a message.
    //the := operator is a shortcut for declaring and initializing a variable in one line
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// init sets initial values for variables used in the function.
//rand.Seed uses the provided seed value to initialize the default Source to a deterministic state. If Seed is not called, the generator behaves as if seeded by Seed(1)
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    //When declaring a slice, you omit its size in the brackets, like this: []string. This tells Go that the array underlying a slice can be dynamically sized.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}