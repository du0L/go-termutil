// +build !windows,!linux,!darwin,!freebsd,cgo

package isatty

/*
#include <unistd.h>
*/
import "C"

import "os"

func Isatty(file *os.File) bool {
    return int(C.isatty(C.int(file.Fd()))) != 0
}
