package main

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func (an Animal) name() string {
	return an.Name
}

//func main() {
//	b := Bird{}
//	b.Name = "Emu"
//	b.Origin = "Australia"
//	b.SpeedKPH = 48
//	b.CanFly = false
//
//	fmt.Printf("%T", b)
//}
