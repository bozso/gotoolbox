package path

import (
    "testing"
)

const nTries = 1000

func testTempFiles(nTries int) (err error) {
    t, err := NewDefaultTempFiles()
    if err != nil {
        return
    }
    
    for ii := 0; ii < nTries; ii++ {
        vf, err := t.Get()
        if err != nil {
            return err
        }
        
        if err = vf.MustExist(); err != nil {
            return err
        }
        
        t.Put(vf)
    }
    return t.Remove()
}


func TestTempFiles(t *testing.T) {
    if err := testTempFiles(nTries); err != nil {
        t.Fatalf("Error: %s\n", err)
    }
}
