package main

import (
	"github.com/fatih/color"
)

func getString() string {
	return "Hello, World!"
}

func main() {
	color.HiMagenta(getString())
}
