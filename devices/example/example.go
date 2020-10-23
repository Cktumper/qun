package example

import "fmt"

type Example struct{}

func NewExample() *Example {
	return &Example{}
}

//	Start Example Device
//
//	Author(Wind)
func (p *Example) Start() error {
	fmt.Println("Hello world")
	return nil
}

//	End Example Device
//
//	Author(Wind)
func (p *Example) End() error {
	return nil
}
