package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	german  = "German"

	greetingPrefixEng = "Hello "
	greetingPrefixSpa = "Hola "
	greetingPrefixFre = "Bonjour "
	greetingPrefixGer = "Halo "
)

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = greetingPrefixFre
	case spanish:
		prefix = greetingPrefixSpa
	case german:
		prefix = greetingPrefixGer
	default:
		prefix = greetingPrefixEng
	}
	return
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(lang) + name + "!"
}

func main() {
	fmt.Println(Hello("Maz", ""))
}
