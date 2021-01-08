package arknights

import "fmt"

type class string

const (
	// Vanguard is "先鋒".
	Vanguard class = "Vanguard"
	// Guard is "前衛".
	Guard class = "Guard"
	// Sniper is "狙撃".
	Sniper class = "Sniper"
	// Specialist is "特殊".
	Specialist class = "Specialist"
	// Caster is "術師".
	Caster class = "Caster"
	// Defender is "重装".
	Defender class = "Defender"
	// Medic is "医療".
	Medic class = "Medic"
)

type operator struct {
	rarity int
	class  class
	elite1 Materials
	elite2 Materials
}

func find(name string) operator {
	m := map[string]operator{
		// ★★★★★★
		"Schwarz": {
			rarity: 6,
			class:  Sniper,
			elite1: Materials{
				Polyester: 8,
				Sugar:     6,
			},
			elite2: Materials{
				D32Steel:    4,
				OrironBlock: 5,
			},
		},
		"Siege": {
			rarity: 6,
			class:  Vanguard,
			elite1: Materials{
				Sugar:  9,
				Device: 3,
			},
			elite2: Materials{
				BipolarNanoflake:     4,
				OrirockConcentration: 6,
			},
		},
		"W": {
			rarity: 6,
			class:  Sniper,
			elite1: Materials{
				OrirockCube: 12,
				Sugar:       5,
			},
			elite2: Materials{
				BipolarNanoflake: 4,
				KetonColloid:     7,
			},
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
