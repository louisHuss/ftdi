package ftdi

/*
#cgo android CFLAGS: -I./libftdi/src -I./libusb/libusb
#cgo android LDFLAGS: ${SRCDIR}/prebuild/android/arm64-v8a/release/libusb.a
#cgo android LDFLAGS: ${SRCDIR}/prebuild/android/arm64-v8a/release/libftdi1.a

#ifdef OS_ANDROID
	#include "os/threads_posix.h"
#endif
*/
import "C"
