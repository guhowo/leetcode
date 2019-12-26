package _055

import "testing"

func Test_CanJump(t *testing.T) {
	for _, unit := range []struct {
		Nums   []int
		Expect bool
	}{
		{
			[]int{2, 3, 1, 1, 4},
			true,
		}, {
			[]int{3, 2, 1, 0, 4},
			false,
		}, {
			[]int{2, 0},
			true,
		}, {
			[]int{2, 0, 0},
			true,
		},
	} {
		if actual := canJump(unit.Nums); actual != unit.Expect {
			t.Errorf("nums = %v, expect = %v, actual = %v", unit.Nums, unit.Expect, actual)
		}
	}
}
