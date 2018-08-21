package main

import "fmt"

// 定义结构体
type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

//func NewSaiyan(name string, power int) Saiyan {
//	return Saiyan{
//		Name: name.
//		Power: power,
//	}
//}

func Super(s *Saiyan) {
	s.Power += 10000
}

func (s *Saiyan) Big() {
	s.Power += 50
}

func main() {
	//	goku := &Saiyan{
	//		Name:  "goku",
	//		Power: 9000,
	//	}
	gohan := &Saiyan{
		Name:  "gohan",
		Power: 5000,
		Father: &Saiyan{
			Name:   "goku",
			Power:  9000,
			Father: nil,
		},
	}
	Super(gohan.Father)
	gohan.Father.Big()
	fmt.Println(gohan.Father.Power)
	fmt.Println(gohan.Father)
}
