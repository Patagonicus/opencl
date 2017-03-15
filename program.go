package opencl

/*
#ifdef __APPLE__
#include <OpenCL/opencl.h>
#else
#include <CL/cl.h>
#endif

#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Program struct {
	program *C.cl_program
}

func (p Program) Release() error {
	err := C.clReleaseProgram(*p.program)
	if err != C.CL_SUCCESS {
		return fmt.Errorf("failed to release program: %d", err)
	}
	C.free(unsafe.Pointer(p.program))
	return nil
}

func (p Program) Build(devices []*Device, options string) error {
	var idPtr *C.cl_device_id
	if len(devices) > 0 {
		idPtr = &asCLDeviceIDs(devices)[0]
	}
	clOptions := C.CString(options)
	defer C.free(unsafe.Pointer(clOptions))
	err := C.clBuildProgram(*p.program, C.cl_uint(len(devices)), idPtr, clOptions, nil, nil)
	if err != C.CL_SUCCESS {
		return fmt.Errorf("failed to build program: %d", err)
	}
	return nil
}

func (p Program) BuildLog(device Device) (string, error) {
	var n C.size_t
	err := C.clGetProgramBuildInfo(*p.program, device.id, C.CL_PROGRAM_BUILD_LOG, C.size_t(0), nil, &n)
	if err != C.CL_SUCCESS {
		return "", fmt.Errorf("failed to get build log size: %d", err)
	}

	result := make([]C.char, n)
	err = C.clGetProgramBuildInfo(*p.program, device.id, C.CL_PROGRAM_BUILD_LOG, n, unsafe.Pointer(&result[0]), nil)
	if err != C.CL_SUCCESS {
		return "", fmt.Errorf("failed to get build log: %d", err)
	}

	return C.GoString(&result[0]), nil
}

func (p Program) CreateKernel(name string) (*Kernel, error) {
	return createKernel(p, name)
}
