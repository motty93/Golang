package model

type Messi struct {
	FootBallPlayer
	SUI int
}

func (m Messi) KickBall() int {
	return m.Power + m.Stamina*m.SUI
}

func (m Messi) Name() string {
	return "Messi"
}
