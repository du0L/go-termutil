// +build windows

package isatty

import (
    "os"
    "syscall"
    "unsafe"
)

var (
    kernel32        = syscall.MustLoadDLL("kernel32.dll")
    fGetConsoleMode = kernel32.MustFindProc("GetConsoleMode")
)

func Isatty(file *os.File) bool {
    var x uint32
    return getConsoleMode(syscall.Handle(file.Fd()), &x) == nil
}

func getConsoleMode(hConsoleHandle syscall.Handle, lpMode *uint32) error {
    ret, _, err := syscall.Syscall(fGetConsoleMode.Addr(), 2,
        uintptr(hConsoleHandle),
        uintptr(unsafe.Pointer(lpMode)),
        0)

    if int(ret) == 0 {
        if err != 0 {
            return error(err)
        } else {
            return syscall.EINVAL
        }
    }
    return
}
