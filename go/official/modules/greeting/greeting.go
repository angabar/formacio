package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("you must provide a name")
	}

	message := fmt.Sprintf(randomString(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, errors.New("you must provide a name")
		}

		messages[name] = message
	}

	return messages, nil
}

func randomString() string {
	greetings := []string{
		"Hi! %v",
		"Hola! %v",
		"Bonjour! %v",
	}

	return greetings[rand.Intn(len(greetings))]
}
