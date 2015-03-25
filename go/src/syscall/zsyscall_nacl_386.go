// mksyscall.pl -l32 -nacl syscall_nacl.go syscall_nacl_386.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package syscall

import "unsafe"

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func naclClose(fd int) (err error) {
	_, _, e1 := Syscall(sys_close, uintptr(fd), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Exit(code int) (err error) {
	_, _, e1 := Syscall(sys_exit, uintptr(code), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func naclFstat(fd int, stat *Stat_t) (err error) {
	_, _, e1 := Syscall(sys_fstat, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func naclRead(fd int, b []byte) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall(sys_read, uintptr(fd), uintptr(_p0), uintptr(len(b)))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func naclSeek(fd int, off *int64, whence int) (err error) {
	_, _, e1 := Syscall(sys_lseek, uintptr(fd), uintptr(unsafe.Pointer(off)), uintptr(whence))
	if e1 != 0 {
		err = e1
	}
	return
}
