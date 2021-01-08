package arknights

import (
	"fmt"
	"sort"
)

type soctier string

const (
	// Low is "初級".
	Low soctier = "low"
	// Mid is "中級".
	Mid soctier = "mid"
)

// Chip is "SoC".
type Chip struct {
	Class  class
	Tier   soctier
	Amount int
}

// Chips is a list of Chip.
type Chips []*Chip

// CountChips returns Chips required for the plan.
func CountChips(plans ...Plan) Chips {
	var res Chips

	add := func(x Chip) {
		for _, v := range res {
			if v.Class == x.Class && v.Tier == x.Tier {
				v.Amount += x.Amount
				return
			}
		}

		res = append(res, &x)
	}

	elite2 := func(rarity int) int {
		switch rarity {
		case 6:
			return 4 * 2
		case 5:
			return 3 * 2
		case 4:
			return 5
		}
		panic(fmt.Sprintf("invalid rarity %v", rarity))
	}

	elite1 := func(rarity int) int {
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

	for _, plan := range plans {
		op := find(plan.Name)

		if plan.Promotion < 2 {
			add(Chip{
				Class:  op.class,
				Tier:   Mid,
				Amount: elite2(op.rarity),
			})
		}

		if plan.Promotion < 1 {
			add(Chip{
				Class:  op.class,
				Tier:   Low,
				Amount: elite1(op.rarity),
			})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].Class < res[j].Class {
			return true
		}
		if res[i].Class > res[j].Class {
			return false
		}
		return res[i].Tier < res[j].Tier
	})

	return res
}
