package main

import (
	"fmt"
	"strings"
)

const (
	spanish = "spanish"
	french  = "french"
	german  = "german"

	engHelloPrefix = "Hello, "
	spaHelloPrefix = "Hola, "
	fraHelloPrefix = "Bonjour, "
	gerHelloPrefix = "Hallo, "
) // wtf u can do this!?

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch strings.ToLower(lang) {
	case spanish:
		prefix = spaHelloPrefix
	case french:
		prefix = fraHelloPrefix
	case german:
		prefix = gerHelloPrefix
	default:
		prefix = engHelloPrefix
	}
	return prefix // you can just write return if u have named the return value e.g (prefix string)

	// if strings.ToLower(lang) == spanish {
	// 	return spaHelloPrefix + name
	// }
	// if strings.ToLower(lang) == french {
	// 	return fraHelloPrefix + name
	// }
}
func main() {
	fmt.Println(Hello("world", ""))
}
