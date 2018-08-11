package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		x, y, want int
	}{
		{100, 50, 150},
		{-100, 50, -50},
		{100, -50, 50},
		{-50, -100, -150},
		{0, 0, 0},
	}
	for _, c := range cases {
		got := Add(c.x, c.y)
		if got != c.want {
			t.Errorf("Add(%#v, %#v) = %#v want %#v", c.x, c.y, got, c.want)
		}
	}
}
