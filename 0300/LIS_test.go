package _300

import "testing"

func Test_LengthOfLIS(t *testing.T) {
	for _, unit := range []struct {
		Nums   []int
		Expect int
	}{
		{
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		}, {
			[]int{10, 9},
			1,
		}, {
			[]int{10},
			1,
		}, {
			[]int{},
			0,
		}, {
			[]int{4, 10, 4, 3, 8, 9},
			3,
		}, {
			[]int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			6,
		},
	} {
		if actual := lengthOfLIS(unit.Nums); actual != unit.Expect {
			t.Errorf("nums = %v, expect = %d, actual = %d", unit.Nums, unit.Expect, actual)
		}
	}
}
