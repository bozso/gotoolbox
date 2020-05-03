package encode

type Encoder interface {
    Encode([]byte, []byte)
    EncodeToString([]byte) string
}

type FileEncoder interface {
    Encode(path string) string
}
