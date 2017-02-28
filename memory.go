package opencl

/*
#ifdef __APPLE__
#include <OpenCL/opencl.h>
#else
#include <CL/cl.h>
#endif
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type MemoryFlags int

const (
	MemReadWrite     = MemoryFlags(C.CL_MEM_READ_WRITE)
	MemWriteOnly     = MemoryFlags(C.CL_MEM_WRITE_ONLY)
	MemReadOnly      = MemoryFlags(C.CL_MEM_READ_ONLY)
	MemUseHostPtr    = MemoryFlags(C.CL_MEM_USE_HOST_PTR)
	MemAllocHostPtr  = MemoryFlags(C.CL_MEM_ALLOC_HOST_PTR)
	MemCopyHostPtr   = MemoryFlags(C.CL_MEM_COPY_HOST_PTR)
	MemHostWriteOnly = MemoryFlags(C.CL_MEM_HOST_WRITE_ONLY)
	MemHostReadOnly  = MemoryFlags(C.CL_MEM_READ_ONLY)
	MemHostNoAccess  = MemoryFlags(C.CL_MEM_HOST_NO_ACCESS)
)

func (f MemoryFlags) asClFlags() C.cl_mem_flags {
	return C.cl_mem_flags(f)
}

type Memory struct {
	memory C.cl_mem
}

func createBuffer(context Context, flags MemoryFlags, size uintptr, hostPtr unsafe.Pointer) (*Memory, error) {
	var err C.cl_int
	memory := C.clCreateBuffer(context.context, flags.asClFlags(), C.size_t(size), hostPtr, &err)
	if err != C.CL_SUCCESS {
		return nil, fmt.Errorf("failed to create buffer: %d", err)
	}
	return &Memory{memory}, nil
}

func (m Memory) Release() error {
	err := C.clReleaseMemObject(m.memory)
	if err != C.CL_SUCCESS {
		return fmt.Errorf("failed to release memory object: %d", err)
	}
	return nil
}
