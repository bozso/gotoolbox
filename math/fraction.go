package math

import (
    "fmt"
    "encoding/json"
)

type Fraction float64

func (f Fraction) String() (s string) {
    return fmt.Sprintf("%f", float64(f))
}

func (f *Fraction) Validate() (err error) {
    return ValidateFraction(float64(*f))
}

func (f *Fraction) UnmarshalJSON(b []byte) (err error) {
    val := float64(0.0)
    
    if err = json.Unmarshal(b, &val); err != nil {
        return
    }
    
    if err = ValidateFraction(val); err != nil {
        return
    }
    
    *f = Fraction(val)
    return
}
