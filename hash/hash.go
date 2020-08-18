package hash

import (
    "math"
    stdHash "hash"
)

// alias of standard library Hash interface
type Hash stdHash.Hash
type Hash32 stdHash.Hash32
type Hash64 stdHash.Hash64

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

type Hasher64 struct {
    stdHash.Hash64
}

func NewHasher64(hash64 stdHash.Hash64) (h Hasher64) {
    h.Hash64 = hash64
    return
}

func (h *Hasher64) Hash(hsb Hashable) {
    h.Reset()
    h.Append(hsb)
}

func (h *Hasher64) Append(hsb Hashable) {
    hsb.Hash(h)
}

func (h *Hasher64) CalcHash(hsb Hashable) (b []byte) {
    h.Hash(hsb)
    return h.Sum(nil)
}

func (h *Hasher64) CalcID(hsb Hashable) (id ID64) {
    h.Hash(hsb)
    return ID64(h.Sum64())
}
