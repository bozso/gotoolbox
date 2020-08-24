package math

import (
    "fmt"
)

/*
 * Helper structure for validating float values. It can be used to
 * check whether a floating point value is in a given range.
 */
type FloatLimiter struct {
    min, max float64
}

func NewFloatLimiter(min, max float64) (f FloatLimiter) {
    f.min, f.max = min, max
    return
}

func (f FloatLimiter) InRangeErr(val float64) (err error) {
    if !f.InRange() {
        err = NotInRange{f:f, val: val}
    }
    return
}

func (f FloatLimiter) InRange(val float64) (b bool) {
    if f.min <= val && val <= f.max {
        b = true
    }
    
    return
}

func (f FloatLimiter) Limit(val float64) (fl float64) {
    if val <= f.min {
        fl = f.min
    } else if val >= f.max {
        fl = f.max
    } else {
        fl = val
    }
    return
}

func (f FloatLimiter) LimitTo(val *float64) {
    *val = f.Limit(*val)
}

var (
    fractionLimiter = NewFloatLimiter(0.0, 1.0)
    percentLimiter = NewFloatLimiter(0.0, 100.0)
)

func ValidatePercent(val float64) (err error) {
    return percentLimiter.InRange(val)
}

func ValidateFraction(val float64) (err error) {
    return fractionLimiter.InRange(val)
}

type NotInRange struct {
    f FloatLimiter
    val float64
}

func (e NotInRange) Error() (s string) {
    return fmt.Sprintf("value '%f' is not in the range of %f - %f",
        e.val, e.f.min, e.f.max)
}
