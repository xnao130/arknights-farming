package arknights

import "testing"

func TestOperators(t *testing.T) {
	for name, op := range Operators {
		if len(op.skill) != 6 {
			t.Errorf("%v: invalid skill l7: %v", name, op.skill)
		}

		if len(op.s1) != 3 {
			t.Errorf("%v: invalid s1: %v", name, op.s1)
		}
		if len(op.s2) != 3 {
			t.Errorf("%v: invalid s2: %v", name, op.s2)
		}

		switch op.rarity {
		case 6:
			if len(op.s3) != 3 {
				t.Errorf("%v: invalid s3: %v", name, op.s3)
			}

		case 4, 5:
			if name == "Amiya" {
				if len(op.s3) != 3 {
					t.Errorf("%v: invalid s3: %v", name, op.s3)
				}
				break
			}

			if len(op.s3) != 0 {
				t.Errorf("%v: invalid s3: %v", name, op.s3)
			}
		}
	}
}
