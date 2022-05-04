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
	prefix := englishHelloPrefix
	switch language {
	case spanish:
		prefix = spanisHelloPrefix
	case french:
		prefix = frechHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
