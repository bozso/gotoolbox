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

func (f FloatLimiter) InRange(val float64) (err error) {
    if f.min <= val && val <= f.max {
        return nil
    }
    
    return NotInRange{f:f, val: val}
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
