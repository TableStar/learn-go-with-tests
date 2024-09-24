package main

import "fmt"

const engHelloPrefix = "Hello, "

func Hello(name string) string {
	// return fmt.Sprintf("Hello, %s", name)
	if name == "" {
		name = "World"
	}
	return engHelloPrefix + name
}
func main() {
	fmt.Println(Hello("world"))
}
