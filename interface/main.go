package main

import (
	"fmt"
	"math"
)

type Monkey struct {
	name   string
	lover  string
	status *Status
}

type Status struct {
	hobby1, hobby2, hobby3, hobby4, hobby5 string
}

func (m *Monkey) getThis() {
	monkey := &Monkey{
		name: "程序员",
		status: &Status{
			hobby1: "格子衫",
			hobby2: "油头",
			hobby3: "人字拖",
			hobby4: "MEVIUS",
			hobby5: "改不完的BUG",
		},
		lover: "纸片人",
	}

	if monkey.lover != "纸片人" {
		fmt.Println("这一定是个假的程序员")
	} else {
		fmt.Println("okay,thanks for your attention")
	}
}

type geometry interface {
	area() float32
	perim() float32
}

type rect struct {
	height, width float32
}

func (r *rect) area() float32 {
	return r.height * r.width
}

func (r *rect) perim() float32 {
	return 2 * (r.height + r.width)
}

type circle struct {
	radius float32
}

func (c *circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float32 {
	return 2 * math.Pi * c.radius
}

func getValue(name string, params interface{}) {
	switch params.(type) {
	case geometry:
		fmt.Println("area of ", name, " is ", params.(geometry).area())
		fmt.Println("perim of ", name, " is ", params.(geometry).perim())
	default:
		fmt.Println("wrong!!")
	}
}

func main() {
	r := &rect{
		height: 2,
		width:  3,
	}
	getValue("rect", r)
	c := &circle{
		radius: 1,
	}
	getValue("circle", c)

	getValue("fuck", "fuck")
}
