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

// A map that describes whether a specific file is in use or not
type TempFileSet map[*ValidFile]bool

// The main struct for managing temporary files in a single directory.
type TempFiles struct {
    // Path to the root directory
    RootDir Dir
    // file set pointing to existing files
    files TempFileSet
    // random number generator for generating random file names
    rand.Rand
}

/*
Set up temporary file management for the specified directory
with the given randum number generator.
*/
func TempFilesFromDir(rootDir Dir, rng rand.Rand) (t TempFiles) {
    return TempFiles{
        RootDir: rootDir,
        files: make(TempFileSet),
        Rand: rng,
    }
}

// Set up temporary file management with default parameters. Random
// number generator will be produced using the current unix time stamp.
func NewDefaultTempFiles() (t TempFiles, err error) {
    src := rand.NewSource(time.Now().Unix())
    t, err = TempFilesFromRand(rand.NoScale(src))
    return
}

// Set up temporary file management with the specified random number
// generator. Directory name will be randomly generated using the
// generator.
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

// Convert it to mutex guarded temporary file manager.
func (t TempFiles) Mutexed() (m MutexTempFiles) {
    return MutexTempFiles{
        tempFiles: t,
    }
}

/*
Search for a valid file that is not in use managed by the receiver.
The second return argument marks whether a file that is not in use was
found.
*/
func (t *TempFiles) Search() (vf *ValidFile, found bool) {
    for file, inUse := range t.files {
        if !inUse {
            vf, t.files[file], found = file, InUse, true
            break
        }
    }
    
    return
}

/*
Retreives a new temporary file to be used.
First it searches for a file that is not in use. If no such file is
found a new file will be created and registered in the receivers
fileset.
*/
func (t *TempFiles) Get() (vf *ValidFile, err error) {
    vf, found := t.Search()
    
    if found {
        return
    }
    
    vf, err = t.NewFile()
    return
}

/*
Creates a new file to be used in the temporary file directory. Returns
error if file creation has failed.
*/
func (t *TempFiles) NewFile() (vf *ValidFile, err error) {
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

/*
Signals to the receiver that the temporary file is no longer in use.
Should be used in conjunction with Get.

    var t = NewDefaultTempFiles()
    f, err := t.Get()
    if err != nil {
        // error handling
    }
    defer t.Put(f)
    // use f
*/
func (t *TempFiles) Put(vf *ValidFile) {
    t.files[vf] = NotInUse
}

/*
Removes the temporary directory containing the temporary files.
*/
func (t *TempFiles) Remove() (err error) {
    err = t.RootDir.Remove()
    return
}

// Error describing file creation failure.
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

// Concurrent safe TempFiles, guarded by mutex
type MutexTempFiles struct {
    // The wrapped struct.
    tempFiles TempFiles
    // mutex for protecting the locking the set
    mutex sync.Mutex
}

// Concurrent safe Get
func (m *MutexTempFiles) Get() (vf *ValidFile, err error) {
    m.mutex.Lock()
    vf, err = m.tempFiles.Get()
    m.mutex.Unlock()
    return
}

// Concurrent safe Search
func (m *MutexTempFiles) Search() (vf *ValidFile, found bool) {
    m.mutex.Lock()
    vf, found = m.tempFiles.Search()
    m.mutex.Unlock()
    return
}

// Concurrent safe NewFile
func (m *MutexTempFiles) NewFile() (vf *ValidFile, err error) {
    m.mutex.Lock()
    vf, err = m.tempFiles.NewFile()
    m.mutex.Unlock()
    return
}

// Concurrent safe Put
func (m *MutexTempFiles) Put(vf *ValidFile) {
    m.mutex.Lock()
    m.tempFiles.Put(vf)
    m.mutex.Unlock()
    return
}
