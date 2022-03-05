package main

import "fmt"

const englishHelloPrefix = "Hello, "
const turkishHelloPrefix = "Maraba, "

func Hello(name string, language string) string {
	if name == "" {
		name = "Matrix"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Turkish":
		prefix = turkishHelloPrefix

	default:
		prefix = englishHelloPrefix

	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))

}
