package main

import (
	"fmt"
	"math/rand"

	"github.com/Golang/youtube/interfaces/model"
)

type Player interface {
	KickBall() int
	Name() string
}

func main() {
	fbp := model.FootBallPlayer{}
	fbp.Stamina = 10
	fbp.Power = 10
	team := make([]Player, 11)
	for i := 0; i < len(team)-2; i++ {
		team[i] = model.FootBallPlayer{
			Stamina: rand.Intn(100),
			Power:   rand.Intn(100),
		}
	}
	team[len(team)-1] = model.CR7{
		FootBallPlayer: model.FootBallPlayer{
			Stamina: rand.Intn(100),
			Power:   rand.Intn(100),
		},
		SUI: rand.Intn(100),
	}
	team[len(team)-2] = model.Messi{
		FootBallPlayer: model.FootBallPlayer{
			Stamina: rand.Intn(100),
			Power:   rand.Intn(100),
		},
		SUI: rand.Intn(100),
	}
	for i := 0; i < len(team); i++ {
		team[i].KickBall()
		fmt.Printf("%s: %d\n", team[i].Name(), team[i].KickBall())
	}
}
