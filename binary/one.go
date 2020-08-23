package binary

type OneByter interface {
    Serializer
    AsByte() (b byte)
}

type Bool struct {
    val bool
    OneByte
}

func (b Bool) AsByte() (b byte) {
    if b.val {
        b = 1
    } else {
        b = 0
    }
    return
}

type Int8 struct {
    val int8
    OneByte
}

func (i Int8) WriteTo(b []byte) {
    
}

int8, uint8, *bool, *int8, *uint8
