package interfaces

type car interface {
	accelerate()
	brake()
}

type ferrari struct {
	model string
	speed int
}

func (f *ferrari) accelerate() {
	f.speed += 10
}

func (f *ferrari) brake() {
	f.speed -= 10
}
