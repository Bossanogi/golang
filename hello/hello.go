package main

import "fmt"

func main() {
	fmt.Println(Hello("word!"))
}

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name + "!"
}
