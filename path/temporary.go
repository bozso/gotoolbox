package path

import (
    "fmt"
    "sync"
    "time"

    "github.com/bozso/emath/rand"
)

const (
    InUse bool = true
    NotInUse bool = false
)

type TempFileSet map[*ValidFile]bool

type TempFiles struct {
    RootDir Dir
    files TempFileSet
    mutex sync.RWMutex
    rand.Rand
}

func TempFilesFromDir(rootDir Dir, rng rand.Rand) (t TempFiles) {
    t.files, t.RootDir, t.Rand = make(TempFileSet), rootDir, rng
    return
}

func NewDefaultTempFiles() (t TempFiles, err error) {
    src := rand.NewSource(time.Now().Unix())
    t, err = TempFilesFromRand(rand.NoScale(src))
    return
}

func TempFilesFromRand(rng rand.Rand) (t TempFiles, err error) {
    prefix := fmt.Sprintf("%d", rng.Int())
    
    t, err = NewTempFiles("", prefix, rng)
    return
}

func NewTempFiles(dir, prefix string, rng rand.Rand) (t TempFiles, err error) {
    d, err := TempDirIn(dir, prefix)
    if err != nil {
        return
    }
    
    t = TempFilesFromDir(d, rng)
    return
}

func (t *TempFiles) Get() (vf *ValidFile, err error) {
    t.mutex.RLock()
    for file, inUse := range t.files {
        if !inUse {
            vf, t.files[file] = file, InUse
            t.mutex.RUnlock()
            return
        }
    }
    
    vf, err = t.NewFile()
    return
}

type CreateFail struct {
    filePath string 
    err error
}

func (e CreateFail) Error() (s string) {
    s = fmt.Sprintf("failed to create temporary file '%s'", e.filePath)
    return
}

func (e CreateFail) Unwrap() (err error) {
    return e.err
}

func (t *TempFiles) NewFile() (vf *ValidFile, err error) {
    t.mutex.Lock()
    defer t.mutex.Unlock()
    
    file := t.RootDir.Join(fmt.Sprintf("%d", t.Rand.Int()))
    
    _, err = file.Create()
    
    if err != nil {
        err = CreateFail{filePath: file.String(), err: err}
        return
    }
    
    vfile, err := file.ToValidFile()
    if err != nil {
        return
    }
    
    vf = &vfile
    
    return
}

func (t *TempFiles) Put(vf *ValidFile) {
    t.files[vf] = NotInUse
}

func (t *TempFiles) Remove() (err error) {
    err = t.RootDir.Remove()
    return
}
