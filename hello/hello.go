package main

import (
	"fmt"
	"strings"
)

const spanish = "spanish"
const french = "french"
const german = "german"
const engHelloPrefix = "Hello, "
const spaHelloPrefix = "Hola, "
const fraHelloPrefix = "Bonjour, "
const gerHelloPrefix = "Hallo, "

func Hello(name, lang string) string {
	// return fmt.Sprintf("Hello, %s", name)
	if name == "" {
		name = "World"
	}
	prefix := engHelloPrefix

	switch strings.ToLower(lang) {
	case spanish:
		prefix = spaHelloPrefix
	case french:
		prefix = fraHelloPrefix
	case german:
		prefix = gerHelloPrefix
	}
	// if strings.ToLower(lang) == spanish {
	// 	return spaHelloPrefix + name
	// }
	// if strings.ToLower(lang) == french {
	// 	return fraHelloPrefix + name
	// }
	return prefix + name
}
func main() {
	fmt.Println(Hello("world", ""))
}
