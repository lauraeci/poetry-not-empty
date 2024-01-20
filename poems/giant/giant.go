package giant

import (
	"fmt"
)

type Giant interface{
	Stomp(aBug string)
	Scream(loudness int)
}

type aGiant struct {
	leftFoot foot
	rightFoot foot 
}

type foot struct {
	furLength int
	toes []toe
	color string
}

type toe struct {
	length int
	color string
}

func (g *aGiant) Stomp(aBug string) {
	fmt.Println("Stomp")
	// I step on a bug and it turns into a bean
	if aBug == "soSmall" {
		fmt.Println("I hardly noticed it")
	} else {
		if aBug == "fliesSoHigh" {
			fmt.Println("I'm so scared!")
			g.Scream(100)
		}
}

// Scream is a method of aGiant that prints a message with the given loudness.
func (g *aGiant) Scream(loudness int) {
	fmt.Println("A giant screams with loudness", loudness)
}


func New() *aGiant {
	me := aGiant{}
	myFeet := []string{"left foot", "right foot"}
	for iLook := 0; iLook < 2; iLook++ {
		fmt.Println(myFeet[iLook])
		me.leftFoot = foot{furLength: 10, color: "brown"}
		me.rightFoot = foot{furLength: 10, color: "brown"}
	}
	return &me
}
