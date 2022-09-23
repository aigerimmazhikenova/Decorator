package main

type Beverage interface {
	cost() float32
	getDescription() string
}
type Coffee struct {
	description string
	Beverage
}
