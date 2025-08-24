package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const japanese = "Japanese"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const japaneseHelloPrefix = "Konichiwa, "
const spanishHelloPrefix = "Hola, " 

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}


