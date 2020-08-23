type ByteSize int

const (
    One ByteSize = 1
    Two          = 2
    Four         = 4
    Eight        = 8
)

type OneByte struct {}

func (_ OneByte) DataSize() (b ByteSize) {
    return One
}

type TwoBytes struct {}

func (_ TwoBytes) DataSize() (b ByteSize) {
    return Two
}

type FourBytes struct {}

func (_ FourBytes) DataSize() (b ByteSize) {
    return Four
}

type EightBytes struct {}

func (_ EightBytes) DataSize() (b ByteSize) {
    return Eight
}
