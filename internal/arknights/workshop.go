package arknights

import "fmt"

func inverse(result material) Materials {
	m := map[material]Materials{}

	for _, class := range []class{Vanguard, Guard, Sniper, Specialist, Caster, Defender, Medic} {
		m[material(fmt.Sprintf("%s Dualchip", class))] = Materials{
			material(fmt.Sprintf("%s Chip Pack", class)): 2,
			ChipCatalyst: 1,
		}
	}

	res, ok := m[result]
	if !ok {
		panic(fmt.Sprintf("invalid material %v", result))
	}
	return res
}

// CountWorkshop returns cost materials to process the workshop plans.
func CountWorkshop(plans Materials) (cost Materials) {
	cost = make(Materials)
	for name, amount := range plans {
		cost = cost.merge(inverse(name).multiply(amount))
	}
	return
}
