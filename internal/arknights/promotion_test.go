package arknights

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCountPromotionMaterials(t *testing.T) {
	testcases := []struct {
		plans    []Plan
		expected Materials
	}{
		{
			plans: []Plan{
				{
					Name:      "Schwarz",
					Promotion: 0,
				},
			},
			expected: Materials{
				D32Steel:    4,
				OrironBlock: 5,
				Polyester:   8,
				Sugar:       6,
			},
		},
		{
			plans: []Plan{
				{
					Name:      "Schwarz",
					Promotion: 1,
				},
			},
			expected: Materials{
				D32Steel:    4,
				OrironBlock: 5,
			},
		},
		{
			plans: []Plan{
				{
					Name:      "Siege",
					Promotion: 1,
				},
				{
					Name:      "W",
					Promotion: 1,
				},
			},
			expected: Materials{
				BipolarNanoflake:     8,
				OrirockConcentration: 6,
				KetonColloid:         7,
			},
		},
	}

	for _, x := range testcases {
		actual := CountPromotionMaterials(x.plans...)

		if diff := cmp.Diff(x.expected, actual); len(diff) != 0 {
			t.Errorf("plans=%v, expected=%v: diff=%v", x.plans, x.expected, diff)
		}
	}
}
