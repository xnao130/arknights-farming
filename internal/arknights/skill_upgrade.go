package arknights

import (
	"fmt"
	"math"
)

func countSkillSummary(level int, rarity int) Materials {
	var name material

	set := func(amount int) Materials {
		return Materials{
			name: amount,
		}
	}

	switch level {
	case 2, 3:
		name = SkillSummary1
		switch rarity {
		case 4:
			return set(2)
		case 5:
			return set(4)
		case 6:
			return set(5)
		}
	case 4, 5, 6:
		name = SkillSummary2
		switch rarity {
		case 4:
			return set(3)
		case 5:
			return set(6)
		case 6:
			return set(8)
		}
	case 7:
		name = SkillSummary3
		switch rarity {
		case 4:
			return set(4)
		case 5:
			return set(6)
		case 6:
			return set(8)
		}
	case 8:
		name = SkillSummary3
		switch rarity {
		case 4:
			return set(2)
		case 5:
			return set(5)
		case 6:
			return set(8)
		}
	case 9:
		name = SkillSummary3
		switch rarity {
		case 4:
			return set(4)
		case 5:
			return set(6)
		case 6:
			return set(12)
		}
	case 10:
		name = SkillSummary3
		switch rarity {
		case 4:
			return set(6)
		case 5:
			return set(10)
		case 6:
			return set(15)
		}
	}
	panic(fmt.Sprintf("invalid level %v or rarity %v", level, rarity))
}

// CountSkillUpgradeMaterials returns skill upgrade materials required for the plan.
func CountSkillUpgradeMaterials(plans ...Plan) Materials {
	res := make(Materials)

	for _, plan := range plans {
		op := find(plan.Name)

		l7 := int(math.Min(math.Max(math.Max(float64(plan.S1), float64(plan.S2)), float64(plan.S3)), 7))
		for i := 1; i < l7; i++ {
			res = res.Add(op.skill[i-1])
			res = res.Add(countSkillSummary(i+1, op.rarity))
		}

		master := func(level int, s []Materials) {
			for i := 0; i < level; i++ {
				res = res.Add(s[i])
				res = res.Add(countSkillSummary(8+i, op.rarity))
			}
		}

		master(plan.S1-7, op.s1)
		master(plan.S2-7, op.s2)
		master(plan.S3-7, op.s3)
	}

	return res
}
