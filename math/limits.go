package math

import (
    "fmt"
)

/*
 * Helper structure for validating float values.
 */
type FloatLimiter struct {
    min, max float64
}

func NewFloatLimiter(min, max float64) (f FloatLimiter) {
    f.min, f.max = min, max
    return
}

// 
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
    return fmt.Sprintf("value '%f' not in range of %f - %f",
        e.val, f.min, f.max)
}
