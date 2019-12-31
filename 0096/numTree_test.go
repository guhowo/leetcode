package _096

import "testing"

func Test_NumTree(t *testing.T) {
	for _, unit := range []struct {
		N      int
		Expect int
	}{
		{
			1,
			1,
		},
		{
			2,
			2,
		},
		{3,
			5,
		},
	} {
		if actual := numTrees(unit.N); actual != unit.Expect {
			t.Errorf("N = %d, Expect = %d, Actual = %d", unit.N, unit.Expect, actual)
		}
	}
}
