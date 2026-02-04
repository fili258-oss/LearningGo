package phrases

import (
	"errors"
	"fmt"
	"math/rand"
)

func GetDay(day string) (string) {

	switch day {
	case "1":
		return "Monday"
	case "2":
		return "Tuesday"
	case "3":
		return "Wednesday"
	case "4":
		return "Thursay"
	case "5":
		return "Friday"
	case "6":
		return "Saturday"
	case "7":
		return "Sunday"
		
	default:
		return "Otro día"
	}

}

func Day(key string) (string, error) {

	if key == "" {
        return key, errors.New("empty key")
    }

	day := GetDay(key)

	message := fmt.Sprintf(randomFormat(), day)

	return message, nil
}

func Days(keys []string) (map[string]string, error){
	messages := make(map[string]string)

	for _, name := range keys{
		message, err := Day(name)
		if err != nil {
			return  nil, err
		}
		daySelected := GetDay(name)
		messages[daySelected] = message

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