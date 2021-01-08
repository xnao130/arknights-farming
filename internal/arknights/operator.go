package arknights

import "fmt"

type class string

const (
	Vanguard   class = "Vanguard"
	Guard      class = "Guard"
	Sniper     class = "Sniper"
	Specialist class = "Specialist"
	Caster     class = "Caster"
	Defender   class = "Defender"
	Medic      class = "Medic"
)

type operator struct {
	rarity int
	class  class
}

func find(name string) operator {
	m := map[string]operator{
		// ★★★★★★
		"Schwarz": {
			rarity: 6,
			class:  Sniper,
		},
		"Siege": {
			rarity: 6,
			class:  Vanguard,
		},
		"W": {
			rarity: 6,
			class:  Sniper,
		},
		"Weedy": {
			rarity: 6,
			class:  Specialist,
		},
		// ★★★★★
		"Amiya": {
			rarity: 5,
			class:  Caster,
		},
		"Bibeak": {
			rarity: 5,
			class:  Guard,
		},
		"Elysium": {
			rarity: 5,
			class:  Vanguard,
		},
		"Lappland": {
			rarity: 5,
			class:  Guard,
		},
		"Liskarm": {
			rarity: 5,
			class:  Defender,
		},
		"Ptilopsis": {
			rarity: 5,
			class:  Medic,
		},
		"Specter": {
			rarity: 5,
			class:  Guard,
		},
		"Texas": {
			rarity: 5,
			class:  Vanguard,
		},
		// ★★★★
		"Dur-nar": {
			rarity: 4,
			class:  Defender,
		},
		"Ethan": {
			rarity: 4,
			class:  Specialist,
		},
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
