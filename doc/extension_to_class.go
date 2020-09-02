package doc

import (
    "fmt"

    "github.com/bozso/gotoolbox/enum"    
)

type (
    typeToExtension map[string]enum.StringSet
)

var class2ext = typeToExtension{
    "text" : enum.NewStringSet("js", "css"),
    "video" : enum.NewStringSet("mp4"),
    "image": enum.NewStringSet("png", "jpg"),
}

func ExtensionToType(ext string) (extType string, err error) {
    for key, val := range class2ext {
        if val.Contains(ext) {
            return key, nil
        }
    }
    
    err = NoMatchingClass{ext}
    return
}

type NoMatchingClass struct {
    Extension string
}

func (e NoMatchingClass) Error() (s string) {
    return fmt.Sprintf("could not find a matching class for extension: '%s'",
        e.Extension)
}

var convert = map[string]string{
    "js": "javascript",
}
