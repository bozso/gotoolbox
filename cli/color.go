package cli

import (
    "fmt"
)

type ColorCode int

const (
    Info ColorCode = iota
    Notice
    Warning
    Error
    Debug
    Bold
)

func Color(s string, color ColorCode) string {
    const (
            info     = "\033[1;34m%s\033[0m"
            notice   = "\033[1;36m%s\033[0m"
            warning  = "\033[1;33m%s\033[0m"
            error_   = "\033[1;31m%s\033[0m"
            debug    = "\033[0;36m%s\033[0m"
            bold     = "\033[1;0m%s\033[0m"
            end      = "\033[0m"
    )
    
    var format string
    
    switch color {
    case Info:
    format = info
    case Notice:
    format = notice
    case Warning:
    format = warning
    case Error:
    format = error_
    case Debug:
    format = debug
    case Bold:
    format = bold
    }
    
    return fmt.Sprintf(format, s)
}

