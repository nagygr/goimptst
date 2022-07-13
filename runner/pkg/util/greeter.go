package util

import (
	"fmt"
)

type Greeter struct {
	Name string
}

func (g *Greeter) Greet() {
	fmt.Printf("Hi %s\n", g.Name)
}
