package arknights

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCountSoC(t *testing.T) {
	actual := CountSoC(plan{
		Name:      "Siege",
		Promotion: 1,
	})

	expected := []soc{
		{
			Class:  Vangard,
			Tier:   Mid,
			Amount: 8,
		},
	}

	if diff := cmp.Diff(expected, actual); len(diff) != 0 {
		t.Errorf("diff=%v", diff)
	}
}
