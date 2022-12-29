package opensimplex

import (
	"math/rand"
	"testing"
	"time"
)

// TestNormNoise_Eval2 verifies that the normalized noise returns values within the expected range
func TestNormNoise_Eval2(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixMilli()))

	for i := 0; i < 100; i++ {
		noise := NewNormalized(rng.Int63())
		// Run one thousand sample points
		for ii := 0; ii < 1000; ii++ {
			x, y := rng.Float64(), rng.Float64()
			if val := noise.Eval2(x, y); val < 0 || val >= 1 {
				t.Fatalf("produced value %.8f", val)
			}
		}
	}
}

func BenchmarkNormNoise_Eval2(b *testing.B) {
	noiseGen := NewNormalized(time.Now().UnixMilli())
	rng := rand.New(rand.NewSource(time.Now().UnixMilli()))
	for i := 0; i < b.N; i++ {
		noiseGen.Eval2(rng.Float64(), rng.Float64())
	}
}
