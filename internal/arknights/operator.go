package arknights

import "fmt"

type class string

const (
	Vangard class = "Vangard"
)

type operator struct {
	rarity int
	class  class
}

func find(name string) operator {
	m := map[string]operator{
		"Siege": {
			rarity: 6,
			class:  Vangard,
		},
	}
	op, ok := m[name]
	if !ok {
		panic(fmt.Sprintf("invalid name %v", name))
	}
	return op
}
