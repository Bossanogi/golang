package main

import "fmt"

func main() {
	helloPrefix := "Spanish"
	fmt.Println(Hello("word!", helloPrefix))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

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
	default:
		helloPrefix = englishHelloPrefix
	}
	return helloPrefix + name + "!"
}
