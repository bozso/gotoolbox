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
//
// TODO(bozso): Implement the mutexed version as a seperate struct?
type TempFiles struct {
    // Path to the root directory
    RootDir Dir
    // file set pointing to existing files
    files TempFileSet
    // mutex for protecting the locking the set
    mutex sync.RWMutex
    // random number generator for generating random file names
    rand.Rand
}

/* Set up temporary file management for the specified directory
 with the given randum number generator. */
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

/*
Search for a valid file that is not in use managed by the receiver.
It locks the receiver's fileset until the end of the operation. The
second return argument marks whether a file that is not in use was
found.
*/
func (t *TempFiles) Search() (vf *ValidFile, found bool) {
    t.mutex.Lock()
    
    for file, inUse := range t.files {
        if !inUse {
            vf, t.files[file], found = file, InUse, true
            break
        }
    }
    
    t.mutex.Unlock()
    return
}

/*
Retreives a new temporary file to be used. Locks the reciever until
the end of the operation. First it searches for a file that is not in
use. If no such file is found a new file will be created and registered
in the receivers fileset.
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
Creates a new file to be used in the temporary file directory. Locks
the receiver until the end of the operation. Returns error if file
creation has failed.
*/
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

/*
Signals to the receiver that the temporary file is no longer in use.
Locks the receiver until the end of the operation. Should be used in
conjunction with Get.

    var t = NewDefaultTempFiles()
    f, err := t.Get()
    if err != nil {
        // error handling
    }
    defer t.Put(f)
    // use f
*/
func (t *TempFiles) Put(vf *ValidFile) {
    t.mutex.Lock()
    t.files[vf] = NotInUse
    t.mutex.Unlock()
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
