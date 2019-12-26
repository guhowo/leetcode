package _053

import "testing"

func Test_MaxSubArray(t *testing.T) {
	for _, unit := range []struct {
		Nums   []int
		Expect int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
	} {
		if act := maxSubArray(unit.Nums); act != unit.Expect {
			t.Errorf("unit = %v, actually = %v, expect = %v", unit, act, unit.Expect)
		}
	}
}
