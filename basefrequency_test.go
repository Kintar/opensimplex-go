package opensimplex

import (
	"testing"
	"time"
)

func TestBaseFrequency(t *testing.T) {
	n := New(time.Now().UnixMilli())
	origin := n.Eval2(0, 0)
	t.Logf("initial sample: %.6f", origin)
	for x := 1.0; x < 1024.0; x++ {
		sample := n.Eval2(x, 0)
		if sample == origin {
			t.Logf("repitition at %.0f", x)
		}
	}
}
