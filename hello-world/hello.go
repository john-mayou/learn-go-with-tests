package main

import (
	"fmt"
	"strings"
)

var langToGreetingPrefix = map[string]string{
	"english": "Hello",
	"spanish": "Hola",
	"french":  "Bonjour",
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	prefix, found := langToGreetingPrefix[strings.ToLower(lang)]
	if !found {
		prefix = "Hello"
	}
	return prefix + ", " + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
