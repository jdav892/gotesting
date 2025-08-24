package main

import "fmt"

const (
		citizen = "Citizen"

		french = "French"
 		spanish = "Spanish"
 		japanese = "Japanese"
		
		earthling = " of Earth"
		america = " from America"
		asia = " from Asia"
		europe = " from Europe"

		france = "France"
		spain = "Spain"
		japan = "Japan"
		usa = "USA"

 		englishHelloPrefix = "Hello, "
 		frenchHelloPrefix = "Bonjour, "
 		japaneseHelloPrefix = "Konichiwa, "
 		spanishHelloPrefix = "Hola, " 
)
func Hello(name string, language string, country string) string {
	if name == "" {
		name = citizen
	}

	return greetingPrefix(language) + name + personalOrigin(country)
}
//public functions have capital letters while private functions have lowercase letters
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

func personalOrigin(country string)(origin string) {
	switch country {
	case spain:
		origin = europe
	case france:
		origin = europe
	case japan:
		origin = asia
	case usa:
		origin = america
	default:
		origin = earthling
	}
	return
}

func main() {
	fmt.Println(Hello("world", "", ""))
}


