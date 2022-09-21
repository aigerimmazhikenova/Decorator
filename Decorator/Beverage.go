package main

type Beverage interface {
	cost() float64
	getDescription() string
}
type Coffee struct {
	description string
	Beverage
}
