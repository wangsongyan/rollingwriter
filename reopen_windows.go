//+build windows

package rollingwriter

import (
	"os"
	"sync/atomic"
	"unsafe"
)

func (w *Writer) reopen(file string) (unsafe.Pointer, error) {
	err := w.file.Close()
	if err != nil {
		return nil, err
	}
	if err := os.Rename(w.absPath, file); err != nil {
		return nil, err
	}
	w.file, err = os.OpenFile(file, DefaultFileFlag, DefaultFileMode)
	if err != nil {
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
