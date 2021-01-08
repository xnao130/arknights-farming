package arknights

import (
	"fmt"
)

type chiptier int

const (
	low chiptier = iota
	mid
	high
)

type chip struct {
	class  class
	tier   chiptier
	amount int
}

func (x chip) materials() Materials {
	var name string
	switch x.tier {
	case low:
		name = "Chip"
	case mid:
		name = "Chip Pack"
	case high:
		name = "Dualchip"
	}

	return Materials{
		material(fmt.Sprintf("%s %s", x.class, name)): x.amount,
	}
}

type chips []*chip

func (x chips) materials() Materials {
	var res Materials
	for _, v := range x {
		res = res.merge(v.materials())
	}
	return res
}

func elite2Amount(rarity int) int {
	switch rarity {
	case 6:
		return 4
	case 5:
		return 3
	case 4:
		return 5
	}
	panic(fmt.Sprintf("invalid rarity %v", rarity))
}

func elite1Amount(rarity int) int {
	switch rarity {
	case 6:
		return 5
	case 5:
		return 4
	case 4:
		return 3
	}
	panic(fmt.Sprintf("invalid rarity %v", rarity))
}

func elite2Tier(rarity int) chiptier {
	switch rarity {
	case 6:
		return high
	case 5:
		return high
	case 4:
		return mid
	}
	panic(fmt.Sprintf("invalid rarity %v", rarity))
}

func countChips(plans ...Plan) chips {
	var res chips

	add := func(x chip) {
		for _, v := range res {
			if v.class == x.class && v.tier == x.tier {
				v.amount += x.amount
				return
			}
		}

		res = append(res, &x)
	}

	for _, plan := range plans {
		op := find(plan.Name)

		if plan.Promotion < 2 {
			add(chip{
				class:  op.class,
				tier:   elite2Tier(op.rarity),
				amount: elite2Amount(op.rarity),
			})
		}

		if plan.Promotion < 1 {
			add(chip{
				class:  op.class,
				tier:   low,
				amount: elite1Amount(op.rarity),
			})
		}
	}

	return res
}
