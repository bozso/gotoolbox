package hash

import (
	"testing"

	"crypto/md5"
	"github.com/bozso/emath/rand"
)

type TestCase struct {
	rng    rand.Rand
	tester Tester
}

func NewTestCase(scale rand.Scale, t Tester) (tc TestCase) {
	tc.rng, tc.tester = scale.New(src), t
	return
}

func (t TestCase) Generate(nfloat int) (floats []float64) {
	floats = make([]float64, nfloat)

	for ii, _ := range floats {
		floats[ii] = t.rng.Float64()
	}

	return
}

func (t TestCase) Test(float float64) (err error) {
	err = t.tester.TestSame(NewFloat64(float))
	return
}

func (t TestCase) TestFloats(floats []float64) (err error) {
	for _, float := range floats {
		if err = t.Test(float); err != nil {
			break
		}
	}
	return
}

const (
	seed = 117
	size = 512
)

var src = rand.NewSource(seed)

func TestFloatHash(t *testing.T) {
	configs := [...]rand.Scale{
		rand.Scale{
			Mean: 12.0,
			Std:  21.0,
		},
		rand.DefaultScale(),
		rand.Scale{
			Mean: -100.0,
			Std:  1.12,
		},
	}

	tester := NewTester(New(md5.New()))

	for _, conf := range configs {
		tc := NewTestCase(conf, tester)

		floats := tc.Generate(size)

		if err := tc.TestFloats(floats); err != nil {
			t.Errorf("%v\n", err)
		}
	}
}
