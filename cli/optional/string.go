package optional

type String struct {
    isSet bool
    s string
}

func (S *String) Set(s string) {
    if len(s) == 0 {
        
    }
}
