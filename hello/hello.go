package main

import "fmt"

func main() {
	helloPrefix := "Spanish"
	fmt.Println(Hello("word!", helloPrefix))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	helloPrefix := englishHelloPrefix
	switch language {
	case "Spanish":
		helloPrefix = spanishHelloPrefix
	case "English":
		helloPrefix = englishHelloPrefix
	case "French":
		helloPrefix = frenchHelloPrefix
	default:
		helloPrefix = englishHelloPrefix
	}
	return helloPrefix + name + "!"
}
