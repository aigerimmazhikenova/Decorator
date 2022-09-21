package main

type DarkRoast struct {
}

func (b *DarkRoast) getDescription() string {
	return "DarkRoast with "
}

func (b *DarkRoast) cost() float64 {
	return 0.99
}
