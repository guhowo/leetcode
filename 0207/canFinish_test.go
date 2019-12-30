package _207

import "testing"

func Test_CanFinish(t *testing.T) {
	for _, unit := range []struct {
		N      int
		Input  [][]int
		Expect bool
	}{
		{
			2,
			[][]int{{1, 0}, {0, 1}},
			false,
		},
		{
			3,
			[][]int{{1, 0}, {1, 2}, {0, 1}},
			false,
		},
		{
			2,
			[][]int{{1, 0}},
			true,
		},
	} {
		if act := canFinish(unit.N, unit.Input); act != unit.Expect {
			t.Errorf("n = %d, input = %v, expect = %v, actual = %v", unit.N, unit.Input, unit.Expect, act)
		}
	}
}
