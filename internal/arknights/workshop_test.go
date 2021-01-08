package arknights

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCountWorkshop(t *testing.T) {
	testcases := []struct {
		plans    Materials
		expected Materials
	}{
		{
			plans: Materials{
				material("Sniper Dualchip"): 1,
			},
			expected: Materials{
				material("Sniper Chip Pack"): 2,
				ChipCatalyst:                 1,
			},
		},
		{
			plans: Materials{
				material("Sniper Dualchip"): 2,
			},
			expected: Materials{
				material("Sniper Chip Pack"): 2 * 2,
				ChipCatalyst:                 1 * 2,
			},
		},
		{
			plans: Materials{
				material("Sniper Dualchip"): 1,
				material("Medic Dualchip"):  2,
			},
			expected: Materials{
				material("Sniper Chip Pack"): 2,
				material("Medic Chip Pack"):  2 * 2,
				ChipCatalyst:                 1 + 1*2,
			},
		},
	}

	for _, x := range testcases {
		actual := CountWorkshop(x.plans)

		if diff := cmp.Diff(x.expected, actual); len(diff) != 0 {
			t.Errorf("plans=%v, expected=%v: diff=%v", x.plans, x.expected, diff)
		}
	}
}
