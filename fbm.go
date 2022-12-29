package opensimplex

// fbmNoise implements Fractional Brownian Motion over a Noise interface
type fbmNoise struct {
	source      Noise
	frequency   float64
	octaves     int
	persistence float64
	lacunarity  float64
}

type evalFunc func(frequency float64) float64

func (f fbmNoise) apply(e evalFunc) float64 {
	sample := 0.0
	for o := 1.0; o <= float64(f.octaves); o++ {
		gain := f.persistence * o
		freq := o * f.lacunarity * f.frequency
		sample += e(freq) * gain
	}
	return sample
}

func (f fbmNoise) Eval2(x, y float64) float64 {
	return f.apply(func(freq float64) float64 {
		return f.source.Eval2(x/freq, y/freq)
	})
}

func (f fbmNoise) Eval3(x, y, z float64) float64 {
	return f.apply(func(freq float64) float64 {
		return f.source.Eval3(x/freq, y/freq, z/freq)
	})
}

func (f fbmNoise) Eval4(x, y, z, w float64) float64 {
	return f.apply(func(freq float64) float64 {
		return f.source.Eval4(x/freq, y/freq, z/freq, w/freq)
	})
}

// NewFbmNoise creates a new FBM Noise function using an input Noise interface.
//
// - octaves are the number of times to apply the noise function
// - frequency is the frequency with which to sample the source noise - higher frequency means faster rate of change
// - lacunarity is a multiplier for frequency applied at each octave
// - persistence is a multiplier for amplitude applied at each octave
//
// **NOTE**: It is not recommended to use
// normalized noise, since application of octaves has the potential to squeeze the output beyond the [0-1) range.
// If you need normalized FBM noise, your options are limited:
//
// - Save your noise data and normalize after generation
// - Run observations on a large data set of points to determine the min and max values, then normalize as you sample
func NewFbmNoise(source Noise, octaves int, frequency, persistence, lacunarity float64) Noise {
	return fbmNoise{
		source:      source,
		frequency:   frequency,
		octaves:     octaves,
		persistence: persistence,
		lacunarity:  lacunarity,
	}
}
