package _09

import "testing"

func Test_Palindarome(t *testing.T) {
	for _, unit := range []struct{
		X int
		Expect bool
	} {
		{
			121,
			true,

		}, {
			-121,
			false,
		},{
			0,
			true,
		}, {
			12345654321,
			true,
		},
	} {
		if actual := isPalindrome(unit.X); actual != unit.Expect {
			t.Errorf("x = %v , expect = %v, actually = %v", unit.X, unit.Expect, actual)
		}

	}
}
