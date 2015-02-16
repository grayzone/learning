package util

import (
    "testing"
)

func TestGetMinutes(t *testing.T) {
	
		cases := []struct {
		in string
		want int
	}{
		{"01:01", 61},
		{"12:01", 721},
		{"", 0},
		{"03:12", 192},
		
	}
	for _, c := range cases {
		got := GetMinutes(c.in)
		if got != c.want {
			t.Errorf("getMinutes(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}

