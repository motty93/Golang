package model

type FootBallPlayer struct {
	Stamina int
	Power   int
}

func (f FootBallPlayer) KickBall() int {
	return f.Power + f.Stamina
}

func (f FootBallPlayer) Name() string {
	return "FootBallPlayer"
}
