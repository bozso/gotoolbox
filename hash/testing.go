package hash

import (
    "fmt"
    "bytes"
)

type HashablePair struct {
    One, Two Hashable
}

func (hp HashablePair) Hash(h Hasher) (op OutPair) {
    op.One, op.Two = h.CalcHash(hp.One), h.CalcHash(hp.Two)
    return
}

func (e HashablePair) Error() (s string) {
    return fmt.Sprintf("got a pair of structures, one (%#v) and two (%#v)",
        e.One, e.Two)
}

type OutPair struct {
    One, Two []byte
}

func (o OutPair) Equals() (b bool) {
    return bytes.Equal(o.One, o.Two)
}

func (e OutPair) Error() (s string) {
    return fmt.Sprintf("got a pair of hash values one '%s' and two '%s'",
        e.One, e.Two)
}

type NotSame struct {
    StructPair HashablePair
    OutPair    OutPair
}

func (e NotSame) Error() (s string) {
    return fmt.Sprintf(
        "identical structures with different hash values, %s, %s",
        e.StructPair.Error(), e.OutPair.Error())
}

type Tester struct {
    hasher Hasher
}

func NewTester(hasher Hasher) (t Tester) {
    t.hasher = hasher
    return
}

func (t Tester) SameHashPair(pair HashablePair) (err error) {
    
    out := pair.Hash(t.hasher)
    
    if !out.Equals() {
        err = NotSame{pair, out}
    }
    return
}

func (t Tester) SameHash(one, two Hashable) (err error) {
    return t.SameHashPair(HashablePair{one, two})
}

func (t Tester) TestSame(hashable Hashable) (err error) {
    err = t.SameHash(hashable, hashable)
    return
}

func (t Tester) TestSames(hashables []Hashable) (err error) {
    for _, hashable := range hashables {
        if err = t.TestSame(hashable); err != nil {
            break
        }
    }
    return
}
