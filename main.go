package main

import (
	"fmt"

	"github.com/lauraeci/poetry-not-empty.git/poems/giant"
)

func main() {
	fmt.Println("Hello World")
	gMe := giant.New()
	gMe.Stomp("soSmall")
	gMe.Stomp("fliesSoHigh")
	gMe.Scream(100)
}