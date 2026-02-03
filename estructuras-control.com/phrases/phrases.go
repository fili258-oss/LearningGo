package phrases

import (
	"errors"
	"fmt"
	"math/rand"
)

func Day(key string) (string, error) {

	if key == "" {
        return key, errors.New("empty key")
    }

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func Days(keys []int) (map[string]string, error){
	messages := make(map[string]string)

	for _, name := range keys{
		message, err := Day(name)
		if err != nil {
			return  nil, err
		}

		messages[name] = message

	}

	return messages, nil
}

func randomFormat() string{

	formats := []string{
        "What a tedious %v!.",
        "It's almost %v",
        "Blessed, %v!",
		"%v, surprise me!",
    }

	return formats[rand.Intn(len(formats))]
}