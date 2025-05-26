package greeter

import "fmt"

type Language string

var phrasebook = map[Language]string{
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"es": "Hola mundo",
	"la": "Salve mundi",
}

func Greet(l Language) (string, error) {
	greeting, ok := phrasebook[l]
	if !ok {
		return "Unsupported language", fmt.Errorf("%q: unsupported language", l)
	}
	return greeting, nil
}
