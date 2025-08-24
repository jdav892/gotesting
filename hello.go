package main

import "fmt"

const (
		french = "French"
 		spanish = "Spanish"
 		japanese = "Japanese"

 		englishHelloPrefix = "Hello, "
 		frenchHelloPrefix = "Bonjour, "
 		japaneseHelloPrefix = "Konichiwa, "
 		spanishHelloPrefix = "Hola, " 
)
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	
	return greetingPrefix(language) + name

}

func greetingPrefix(language string)(prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}


