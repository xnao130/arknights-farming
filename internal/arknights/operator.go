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
			elite1: Materials{
				Device: 6,
				Sugar:  4,
			},
			elite2: Materials{
				D32Steel:            4,
				ManganeseTrihydrate: 6,
			},
		},
		// ★★★★★
		"Amiya": {
			rarity: 5,
			class:  Caster,
			elite1: Materials{
				Device: 4,
				Oriron: 4,
			},
			elite2: Materials{
				OrirockConcentration: 10,
				LoxicKohl:            10,
			},
		},
		"Bibeak": {
			rarity: 5,
			class:  Guard,
			elite1: Materials{
				Oriron:      4,
				OrirockCube: 3,
			},
			elite2: Materials{
				ManganeseTrihydrate: 8,
				RMA7012:             8,
			},
		},
		"Elysium": {
			rarity: 5,
			class:  Vanguard,
			elite1: Materials{
				Polyester: 4,
				Sugar:     3,
			},
			elite2: Materials{
				IncandescentAlloyBlock: 7,
				Aketon:                 16,
			},
		},
		"Lappland": {
			rarity: 5,
			class:  Guard,
			elite1: Materials{
				Device:      3,
				OrirockCube: 4,
			},
			elite2: Materials{
				OptimizedDevice: 6,
				OrironCluster:   10,
			},
		},
		"Liskarm": {
			rarity: 5,
			class:  Defender,
			elite1: Materials{
				Sugar:       5,
				OrirockCube: 3,
			},
			elite2: Materials{
				GrindstonePentahydrate: 7,
				Aketon:                 15,
			},
		},
		"Ptilopsis": {
			rarity: 5,
			class:  Medic,
			elite1: Materials{
				OrirockCube: 8,
				Sugar:       2,
			},
			elite2: Materials{
				OrirockConcentration: 9,
				Grindstone:           10,
			},
		},
		"Specter": {
			rarity: 5,
			class:  Guard,
			elite1: Materials{
				OrirockCube: 6,
				Polyketon:   3,
			},
			elite2: Materials{
				WhiteHorseKohl: 8,
				Aketon:         15,
			},
		},
		"Texas": {
			rarity: 5,
			class:  Vanguard,
			elite1: Materials{
				Polyester: 5,
				Oriron:    3,
			},
			elite2: Materials{
				PolyesterLump:  8,
				OrirockCluster: 16,
			},
		},
		// ★★★★
		"Dur-nar": {
			rarity: 4,
			class:  Defender,
			elite1: Materials{
				OrirockCube: 1,
				Polyester:   1,
			},
			elite2: Materials{
				OrirockCluster: 19,
				RMA7012:        8,
			},
		},
		"Ethan": {
			rarity: 4,
			class:  Specialist,
			elite1: Materials{
				Sugar:  1,
				Oriron: 1,
			},
			elite2: Materials{
				SugarPack:      17,
				OrirockCluster: 14,
			},
		},
		"Utage": {
			rarity: 4,
			class:  Guard,
			elite1: Materials{
				Device: 1,
				Sugar:  1,
			},
			elite2: Materials{
				Aketon:         14,
				OrirockCluster: 14,
			},
		},
	}
	op, ok := m[name]
	if !ok {
		panic(fmt.Sprintf("invalid name %v", name))
	}
	return op
}
