package command

import (

)

type VCS interface {
    Status() (b []byte, err error)
}
