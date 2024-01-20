package model

type CR7 struct {
	FootBallPlayer
	SUI int
}

func (c CR7) KickBall() int {
	return c.Power + c.Stamina*c.SUI
}

func (c CR7) Name() string {
	return "CR7"
}
