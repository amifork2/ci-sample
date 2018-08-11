package sub

import (
	"testing"
)

func TestSub(t *testing.T) {
	cases := []struct {
		x, y, want int
	}{
		{100, 50, 50},
		{-100, 50, -150},
		{100, -50, 150},
		{-50, -100, 50},
		{0, 0, 0},
	}
	for _, c := range cases {
		got := Sub(c.x, c.y)
		if got != c.want {
			t.Errorf("Sub(%#v, %#v) = %#v want %#v", c.x, c.y, got, c.want)
		}
	}
}
