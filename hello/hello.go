package main

import "fmt"

func main() {
	helloPrefix := "Spanish"
	fmt.Println(Hello("word!", helloPrefix))
}

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

const (
	spanish = "Spanish"
	french  = "French"
	english = "English"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name + "!"
}
func greetingPrefix(language string) (helloPrefix string) {
	switch language {
	case spanish:
		helloPrefix = spanishHelloPrefix
	case english:
		helloPrefix = englishHelloPrefix
	case french:
		helloPrefix = frenchHelloPrefix
	default:
		helloPrefix = englishHelloPrefix
	}
	return
}
