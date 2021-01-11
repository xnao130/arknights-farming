package arknights

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCountSkillUpgradeMaterials(t *testing.T) {
	testcases := []struct {
		plans    []Plan
		expected Materials
	}{
		{
			plans: []Plan{
				{
					Name: "Schwarz",
					S1:   1,
				},
			},
			expected: Materials{},
		},
		{
			plans: []Plan{
				{
					Name: "Schwarz",
					S1:   2,
				},
			},
			expected: Materials{
				SkillSummary1: 5,
			},
		},
		{
			plans: []Plan{
				{
					Name: "Schwarz",
					S1:   7,
				},
			},
			expected: Materials{
				SkillSummary1:   10,
				SkillSummary2:   24,
				SkillSummary3:   8,
				SugarSubstitute: 5,
				Diketon:         4,
				Polyester:       5,
				OrironShard:     4,
				Sugar:           3,
				LoxicKohl:       7,
				ManganeseOre:    3,
				RMA7012:         4,
			},
		},
		{
			plans: []Plan{
				{
					Name: "Schwarz",
					S1:   7,
					S2:   7,
					S3:   7,
				},
			},
			expected: Materials{
				SkillSummary1:   10,
				SkillSummary2:   24,
				SkillSummary3:   8,
				SugarSubstitute: 5,
				Diketon:         4,
				Polyester:       5,
				OrironShard:     4,
				Sugar:           3,
				LoxicKohl:       7,
				ManganeseOre:    3,
				RMA7012:         4,
			},
		},
		{
			plans: []Plan{
				{
					Name: "Schwarz",
					S1:   8,
					S2:   9,
					S3:   10,
				},
			},
			expected: Materials{
				SkillSummary1:             10,
				SkillSummary2:             24,
				SkillSummary3:             8 + 8*3 + 12*2 + 15,
				SugarSubstitute:           5,
				Diketon:                   4,
				Polyester:                 5,
				OrironShard:               4,
				Sugar:                     3,
				LoxicKohl:                 7 + 7,
				ManganeseOre:              3 + 5,
				RMA7012:                   4,
				GrindstonePentahydrate:    4,
				RMA7024:                   4,
				PolyesterLump:             4,
				OrirockConcentration:      10 + 4,
				Grindstone:                7,
				OrironBlock:               4,
				SugarLump:                 7,
				PolymerizationPreparation: 6,
				ManganeseTrihydrate:       6,
			},
		},
	}

	for _, x := range testcases {
		actual := CountSkillUpgradeMaterials(x.plans...)

		if diff := cmp.Diff(x.expected, actual); len(diff) != 0 {
			t.Errorf("plans=%v, expected=%v: diff=%v", x.plans, x.expected, diff)
		}
	}
}
