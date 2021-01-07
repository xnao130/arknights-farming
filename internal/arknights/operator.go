package arknights

import "fmt"

type class string

const (
	Vanguard class = "Vanguard"
	Guard    class = "Guard"
)

type operator struct {
	rarity int
	class  class
}

func find(name string) operator {
	m := map[string]operator{
		// ★★★★★★
		"Siege": {
			rarity: 6,
			class:  Vanguard,
		},
		// ★★★★★
		"Elysium": {
			rarity: 5,
			class:  Vanguard,
		},
		// ★★★★
		"Utage": {
			rarity: 4,
			class:  Guard,
		},
	}
	op, ok := m[name]
	if !ok {
		panic(fmt.Sprintf("invalid name %v", name))
	}
	return op
}
