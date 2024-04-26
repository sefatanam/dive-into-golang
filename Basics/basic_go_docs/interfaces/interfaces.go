package interfaces

import (
	"fmt"
	"math/rand"
)

type Striker struct {
	Name  string
	Power int
}

func (s Striker) KickBall() {
	speed := s.Power * rand.Intn(s.Power)
	fmt.Printf("Striker kick the ball at speed %d\n", speed)
}

type Defender struct {
	Name     string
	Strength int
}

func (d Defender) KickBall() {
	speed := d.Strength + rand.Intn(d.Strength)
	fmt.Printf("Defender kick the ball at speed %d\n", speed)
}

type Player interface {
	KickBall()
}

func InterfacesMain() {
	players := []Player{
		Striker{"Wallace", 20},
		Defender{"Ninja", 12},
	}

	for _, player := range players {
		player.KickBall()
	}
}
