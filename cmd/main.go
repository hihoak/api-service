package main

import (
	"github.com/fatih/color"
)

const greetingsMessage = "Hello, World!"

func getString() string {
	return greetingsMessage
}

func main() {
	color.HiMagenta(getString())
}
