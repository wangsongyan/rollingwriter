//+build !windows

package rollingwriter

import (
	"os"
	"sync/atomic"
	"unsafe"
)

func (w *Writer) reopen(file string) (unsafe.Pointer, error) {
	if err := os.Rename(w.absPath, file); err != nil {
		return nil, err
	}
	newfile, err := os.OpenFile(w.absPath, DefaultFileFlag, DefaultFileMode)
	if err != nil {
		return nil, err
	}

	// swap the unsafe pointer
	oldfile := atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&w.file)), unsafe.Pointer(newfile))
	return oldfile, nil
}
