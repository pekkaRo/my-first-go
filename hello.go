package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const frechHelloPrefix = "Bonjour, "
const spanisHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanisHelloPrefix + name
	}
	if language == french {
		return frechHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
