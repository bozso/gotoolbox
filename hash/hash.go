package hash

import (
    "math"
    stdHash "hash"
)

type ID32 uint32
type ID64 uint64

type Hashable interface {
    Hash(h stdHash.Hash)
}

type HashFloat64 struct {
    mantissa uint64
    exponent int16
    sign int8
}

func (hf HashFloat64) Hash(h stdHash.Hash) {
}

func HashF64(val float64, h stdHash.Hash) {
    NewHashFloat64(val).Hash(h)
}

func NewHashFloat64(val float64) (h HashFloat64) {
    bits := math.Float64bits(val)
    sign := int8(-1)
    
    if bits >> 63 == 0 {
        sign = int8(1)
    }

    exponent := int16((bits >> 52) & 0x7ff)
    
    mantissa := (bits & 0xfffffffffffff) | 0x10000000000000
    
    if exponent == 0 {
        mantissa = (bits & 0xfffffffffffff) << 1
    }

    exponent -= 1023 + 52;
    
    return HashFloat64{
        exponent: exponent,
        sign: sign,
        mantissa: mantissa,
    }
}

type Hasher32 struct {
    stdHash.Hash32
}

func NewHasher32(hash32 stdHash.Hash32) (h Hasher32) {
    h.Hash32 = hash32
    return
}

func (h *Hasher32) Hash(hsb Hashable) {
    h.Reset()
    h.Append(hsb)
}

func (h *Hasher32) Append(hsb Hashable) {
    hsb.Hash(h)
}

func (h *Hasher32) CalcHash(hsb Hashable) (b []byte) {
    h.Hash(hsb)
    return h.Sum(nil)
}

func (h *Hasher32) CalcID(hsb Hashable) (id ID32) {
    h.Hash(hsb)
    return ID32(h.Sum32())
}
