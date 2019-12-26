package _026

import "testing"

func Test_RemoveDuplicates(t *testing.T) {
	for _, unit := range []struct {
		Nums   []int
		Expect int
	}{
		{
			[]int{1, 1, 2},
			2,
		}, {
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		}, {
			[]int{1, 1, 1, 2},
			2,
		},
	} {
		if actual := removeDuplicates(unit.Nums); actual != unit.Expect {
			t.Errorf("Nums = %v, expect = %d, actual = %d", unit.Nums, unit.Expect, actual)
		}
	}
}
