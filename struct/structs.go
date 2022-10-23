package main

type Doctor struct {
	number     int
	actorName  string
	episodes   []string
	companions []string
}

//func main() {
//	aDoctor := Doctor{
//		number:    3,
//		actorName: "Jon Pertwee",
//		episodes: []string{
//			"Doctor Who and the Silurians",
//		},
//		companions: []string{
//			"Liz Shaw",
//			"Jo Grant",
//			"Sarah Jane Smith",
//		},
//	}
//
//	//Anonymous struct
//	bDoctor := struct{ name string }{name: "John Pertwee"}
//
//	println(aDoctor.actorName)
//	println(aDoctor.companions[1])
//	println(bDoctor.name)
//}
