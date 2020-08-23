package hash

import (
    "math"
    "encoding/binary"
    stdHash "hash"
)

// alias to the standard library Hash interface
type Hash stdHash.Hash
type Hash32 stdHash.Hash32
type Hash64 stdHash.Hash64

// alias for return types of the Sum functions
type ID32 uint32
type ID64 uint64

// A type that can be hashed using the Hash interface
type Hashable interface {
    Hash(h stdHash.Hash)
}

type Float64 struct {
    mantissa uint64
    exponent int16
    sign int8
}

func (f Float64) Hash(h stdHash.Hash) {
    binary.Write(h, binary.LittleEndian, f.mantissa)
    binary.Write(h, binary.LittleEndian, f.exponent)
    binary.Write(h, binary.LittleEndian, f.sign)
}

func HashFloat64(val float64, h stdHash.Hash) {
    NewFloat64(val).Hash(h)
}

func NewFloat64(val float64) (h Float64) {
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
    
    return Float64{
        exponent: exponent,
        sign: sign,
        mantissa: mantissa,
    }
}

type Hasher struct {
    hasher Hash
}

func New(hasher Hash) (h Hasher) {
    h.hasher = hasher
    return
}

func (h *Hasher) Append(hsb Hashable) {
    hsb.Hash(h.hasher)
}

func (h *Hasher) Hash(hsb Hashable) {
    h.hasher.Reset()
    h.Append(hsb)
}

func (h *Hasher) CalcHash(hsb Hashable) (b []byte) {
    h.Hash(hsb)
    return h.hasher.Sum(nil)
}

type Hasher32 struct {
    Hasher
    stdHash.Hash32
}

func New32(hash32 stdHash.Hash32) (h Hasher32) {
    h.Hasher = New(hash32)
    h.Hash32 = hash32
    return
}

func (h *Hasher32) CalcID(hsb Hashable) (id ID32) {
    h.Hash(hsb)
    return ID32(h.Sum32())
}

type Hasher64 struct {
    Hasher
    stdHash.Hash64
}

func New64(hash64 stdHash.Hash64) (h Hasher64) {
    h.Hasher = New(hash64)
    h.Hash64 = hash64
    return
}

func (h *Hasher64) CalcID(hsb Hashable) (id ID64) {
    h.Hash(hsb)
    return ID64(h.Sum64())
}
