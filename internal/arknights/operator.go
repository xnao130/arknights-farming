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
	skill  []Materials
	s1     []Materials
	s2     []Materials
	s3     []Materials
}

// Operators is a list of operators.
var Operators = map[string]operator{
	// ★★★★★★
	"Schwarz": Schwarz,
	"Siege":   Siege,
	"W":       W,
	"Weedy":   Weedy,
	// ★★★★★
	"Amiya":     Amiya,
	"Bibeak":    Bibeak,
	"Elysium":   Elysium,
	"Lappland":  Lappland,
	"Liskarm":   Liskarm,
	"Ptilopsis": Ptilopsis,
	"Specter":   Specter,
	"Texas":     Texas,
	// ★★★★
	"Dur-nar": DurNar,
	"Ethan":   Ethan,
	"Utage":   Utage,
}

func find(name string) operator {
	op, ok := Operators[name]
	if !ok {
		panic(fmt.Sprintf("invalid name %v", name))
	}
	return op
}
