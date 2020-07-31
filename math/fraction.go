package math

import (
    "encoding/json"
)

type Fraction float64

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
