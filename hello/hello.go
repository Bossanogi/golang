package main

import "fmt"

func main() {
	fmt.Println(Hello("word!"))
}

func Hello(name string) string {
	return "Hello, world!"
}
