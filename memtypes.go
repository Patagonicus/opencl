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
)

type MemCacheType int

const (
	CacheNone = MemCacheType(iota)
	CacheReadOnly
	CacheReadWrite
)

func memCacheType(t C.cl_device_mem_cache_type) (MemCacheType, error) {
	switch t {
	case C.CL_NONE:
		return CacheNone, nil
	case C.CL_READ_ONLY_CACHE:
		return CacheReadOnly, nil
	case C.CL_READ_WRITE_CACHE:
		return CacheReadWrite, nil
	default:
		return CacheNone, fmt.Errorf("unknown mem cache type: %d", t)
	}
}

func (t MemCacheType) Name() string {
	switch t {
	case CacheNone:
		return "none"
	case CacheReadOnly:
		return "read only"
	case CacheReadWrite:
		return "read write"
	default:
		return "unknown"
	}
}

func (t MemCacheType) String() string {
	return fmt.Sprintf("MemCacheType{%s}", t.Name())
}

type LocalMemType int

const (
	LocalMemNone = LocalMemType(iota)
	LocalMemLocal
	LocalMemGlobal
)

func localMemType(t C.cl_device_local_mem_type) (LocalMemType, error) {
	switch t {
	case C.CL_NONE:
		return LocalMemNone, nil
	case C.CL_LOCAL:
		return LocalMemLocal, nil
	case C.CL_GLOBAL:
		return LocalMemGlobal, nil
	default:
		return LocalMemNone, fmt.Errorf("unknown local memory type: %d", t)
	}
}

func (t LocalMemType) Name() string {
	switch t {
	case LocalMemNone:
		return "none"
	case LocalMemLocal:
		return "local"
	case LocalMemGlobal:
		return "global"
	default:
		return "unknown"
	}
}

func (t LocalMemType) String() string {
	return fmt.Sprintf("LocalMemType{%s}", t.Name())
}
