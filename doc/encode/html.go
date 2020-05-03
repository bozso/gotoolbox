package encode

type (
    extensionSet map[string]struct{}
    typeToExtension map[string]extensionSet
)

var class2ext = typeToExtension{
    "text" : extensionSet{
        "js": struct{}, "css": struct{},
    },
    
    "video" : extensionSet{
        "mp4": struct{},
    }),
    
    "image": extensionSet{
        "png": struct{}, "jpg": struct{},
    },
}

var convert = map[string]string{
    "js": "javascript",
}

type HTML struct {
    encoder FileEncoder
}

func NewHtml(encoder FileEncoder) (h HTML) {
    return HTML{
        encoder: encoder,
    }
}

func (h HTML) Encode(path string) (r Result, err error) {
    
}
