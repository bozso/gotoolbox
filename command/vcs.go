package command

import ()

/*
Represents a common set of functions implemented by most version
control systems.

TODO(bozso): Set up functions for other functionalities.
*/
type VCS interface {
	Status() (b []byte, err error)
}
