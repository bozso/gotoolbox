//+build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Install() error {
	return sh.Run(mg.GoCmd(), "install")
}
