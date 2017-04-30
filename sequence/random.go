package sequence

import "math/rand"

type Random struct {
	rnd     *rand.Rand
	scale   *float64
	average *float64
	len     int
}

func (r *Random) Len() int {
	return r.len
}

func (r *Random) GetValue(_ int) float64 {
	if r.scale != nil {
		return r.rnd.Float64() * *r.scale
	}
	return r.rnd.Float64()
}
