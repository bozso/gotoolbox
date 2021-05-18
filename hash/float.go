package hash

import (
	"encoding/binary"
	"hash"
	"math"
)

type Float struct {
	mantissa uint64
	exponent int16
	sign     int8
}

func NewFloat64(val float64) (h Float) {
	bits := math.Float64bits(val)
	sign := int8(-1)

	if bits>>63 == 0 {
		sign = int8(1)
	}

	exponent := int16((bits >> 52) & 0x7ff)

	mantissa := (bits & 0xfffffffffffff) | 0x10000000000000

	if exponent == 0 {
		mantissa = (bits & 0xfffffffffffff) << 1
	}

	exponent -= 1023 + 52

	return Float{
		exponent: exponent,
		sign:     sign,
		mantissa: mantissa,
	}
}

func (f Float) Hash(h hash.Hash) {
	// hash.Hash.Write never returns an error
	_ = binary.Write(h, binary.LittleEndian, f.mantissa)
	_ = binary.Write(h, binary.LittleEndian, f.exponent)
	_ = binary.Write(h, binary.LittleEndian, f.sign)
}

func Float64(val float64, h hash.Hash) {
	NewFloat64(val).Hash(h)
}
