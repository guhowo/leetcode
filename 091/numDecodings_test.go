package _91

import "testing"

func Test_NumDecodings(t *testing.T) {
	for _, unit := range []struct {
		Input  string
		Expect int
	}{
		{
			"110",
			1,
		},
		{
			"0",
			0,
		},
		{
			"101",
			1,
		},
		{
			"12",
			2,
		},
		{
			"404",
			0,
		}, {
			"10",
			1,
		}, {
			"1212",
			5,
		},
	} {
		if act := numDecodings(unit.Input); act != unit.Expect {
			t.Errorf("case = %v, output = %v, expect = %v", unit, act, unit.Expect)
		}
	}
}
