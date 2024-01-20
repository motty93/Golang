package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	KickBall() int
	Name() string
}

type CR7 struct {
	stamina int
	power   int
	SUI     int
}

type Messi struct {
	stamina int
	power   int
	SUI     int
}

type FootBallPlayer struct {
	stamina int
	power   int

	// field
}

func (c CR7) KickBall() int {
	return c.power + c.stamina*c.SUI
}

func (c CR7) Name() string {
	return "CR7"
}

func (m Messi) KickBall() int {
	return m.power + m.stamina*m.SUI
}

func (m Messi) Name() string {
	return "Messi"
}

func (f FootBallPlayer) KickBall() int {
	return f.power + f.stamina
}

func (f FootBallPlayer) Name() string {
	return "FootBallPlayer"
}

func main() {
	team := make([]Player, 11)
	for i := 0; i < len(team)-2; i++ {
		team[i] = FootBallPlayer{
			stamina: rand.Intn(100),
			power:   rand.Intn(100),
		}
	}
	team[len(team)-1] = CR7{
		stamina: rand.Intn(100),
		power:   rand.Intn(100),
		SUI:     rand.Intn(100),
	}
	team[len(team)-2] = Messi{
		stamina: rand.Intn(100),
		power:   rand.Intn(100),
		SUI:     rand.Intn(100),
	}
	for i := 0; i < len(team); i++ {
		team[i].KickBall()
		fmt.Printf("%s: %d\n", team[i].Name(), team[i].KickBall())
	}
}
