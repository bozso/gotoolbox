//+build mage

package main

import (
    "github.com/magefile/mage/sh"
    "github.com/magefile/mage/mg"
)

func Install() error {
    return sh.Run(mg.GoCmd(), "install")
}
