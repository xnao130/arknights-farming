package arknights

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCountSoC(t *testing.T) {
	testcases := []struct {
		plans    []Plan
		expected []*soc
	}{
		{
			plans: []Plan{
				{
					Name:      "Siege",
					Promotion: 1,
				},
			},
			expected: []*soc{
				{
					Class:  Vanguard,
					Tier:   Mid,
					Amount: 8,
				},
			},
		},
		{
			plans: []Plan{
				{
					Name:      "Siege",
					Promotion: 1,
				},
				{
					Name:      "Elysium",
					Promotion: 0,
				},
			},
			expected: []*soc{
				{
					Class:  Vanguard,
					Tier:   Low,
					Amount: 4,
				},
				{
					Class:  Vanguard,
					Tier:   Mid,
					Amount: 14,
				},
			},
		},
		{
			plans: []Plan{
				{
					Name:      "Siege",
					Promotion: 1,
				},
				{
					Name:      "Elysium",
					Promotion: 0,
				},
				{
					Name:      "Utage",
					Promotion: 1,
				},
			},
			expected: []*soc{
				{
					Class:  Guard,
					Tier:   "mid",
					Amount: 5,
				},
				{
					Class:  Vanguard,
					Tier:   Low,
					Amount: 4,
				},
				{
					Class:  Vanguard,
					Tier:   Mid,
					Amount: 14,
				},
			},
		},
	}

	for _, x := range testcases {
		actual := CountSoC(x.plans...)

		if diff := cmp.Diff(x.expected, actual); len(diff) != 0 {
			t.Errorf("plans=%v, expected=%v: diff=%v", x.plans, x.expected, diff)
		}
	}
}
