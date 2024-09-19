package helloworld

import "strings"

var prefix = map[string]string{
	"english": "Hello, ",
	"spanish": "Hola, ",
	"french":  "Bonjour, ",
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "" {
		language = "English"
	}

	greeting := prefix[strings.ToLower(language)]

	return greeting + name
}

func main() {
	Hello("Dave", "English")
}
