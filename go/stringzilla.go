package stringzilla

// #cgo CFLAGS: -I${SRCDIR}/../include -std=c99 -Wno-unknown-pragmas -Wno-uninitialized -Wno-cast-function-type -Wno-unused-function
// #include <stringzilla/stringzilla.h>
import "C"
import "unsafe"

func ToUpper(s string) string {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.sz_toupper(cstr, C.size_t(len(s)), cstr)
	return C.GoString(cstr)
}

func Find(haystack string, needle string) (found bool, index uint64) {
	chaystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(chaystack))
	cneedle := C.CString(needle)
	defer C.free(unsafe.Pointer(cneedle))
	matchPtr := unsafe.Pointer(C.sz_find(chaystack, C.size_t(len(haystack)), cneedle, C.size_t(len(needle))))
	if matchPtr == nil {
		found = false
	} else {
		found = true
		index = uint64(uintptr(matchPtr) - uintptr(unsafe.Pointer(chaystack)))
	}
	return
}
