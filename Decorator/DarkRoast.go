package main

type DarkRoast struct {
}

func (b *DarkRoast) getDescription() string {
	return "Yor order: \n \t DarkRoast   - 0.99$"
}

func (b *DarkRoast) cost() float32 {
	return 0.99
}
