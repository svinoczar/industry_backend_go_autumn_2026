package main

import "strings"

func greet(name string) string {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		name = "World"
	}
	return "Hello, " + name + "!"
}
