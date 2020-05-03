package encode

import (
    "os"
    "bytes"
    "io/ioutil"
    "net/http"
    "regexp"
)

const urlRegexStr = "^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$"
var urlRegex = regexp.MustCompile(urlRegexStr)


type Loader interface {
    Load(string) ([]byte, error)
}

type LoadFromFile struct{}

func (_ LoadFromFile) Load(path string) (b []byte, err error) {
    f, err := os.Open(path)
    if err != nil {
        return
    }
    defer f.Close()
    
    b, err = ioutil.ReadAll(f)
    return
}

type LoadFromUrl struct{}

func (_ LoadFromUrl) Load(url string) (b []byte, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    out := bytes.Buffer{}
    _, err = io.Copy(out, resp.Body)

    b = out.Bytes()
    
    return
}

var (
    fromUrl = LoadFromUrl{}
    fromFile = LoadFromFile{}
)

func LoadMedia(path string) (b []byte, err error) {
    matched, err := urlRegex.Match(path)
    if err != nil {
        return
    }
    
    if matched {
        return fromUrl.Load(path)
    }
    
    return fromFile.Load(path)
}

    def __call__(self, media_path):
        mode = ext = path.splitext(media_path)[1].strip(".")
        
        for key, val in self.ext2klass.items():
            if ext in key:
                klass = val
                break
        
        if mode in self.convert:
            mode = convert[ext]
        
        data = self.encoder(self.load_media(media_path))
        
        return self.tpl.format(
            klass=klass, mode=mode, data=data.decode("utf-8")
        )
