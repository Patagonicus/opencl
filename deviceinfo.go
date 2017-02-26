// Code generated go generate DO NOT EDIT.
package opencl

/*
#include <stdlib.h>

#ifdef __APPLE__
#include <OpenCL/opencl.h>
#else
#include <CL/cl.h>
#endif
*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

func (d Device) getInfoType(name string, id C.cl_device_info) (DeviceType, error) {
	var buf C.cl_device_type
	if err := C.clGetDeviceInfo(d.id, id, C.size_t(unsafe.Sizeof(buf)), unsafe.Pointer(&buf), nil); err != C.CL_SUCCESS {
		return 0, fmt.Errorf("error getting device info %s: %d", name, err)
	}

	return DeviceType(buf), nil
}

func (d Device) getInfoBool(name string, id C.cl_device_info) (bool, error) {
	var buf C.cl_bool
	if err := C.clGetDeviceInfo(d.id, id, C.size_t(unsafe.Sizeof(buf)), unsafe.Pointer(&buf), nil); err != C.CL_SUCCESS {
		return false, fmt.Errorf("error getting device info %s: %d", name, err)
	}

	return buf == C.CL_TRUE, nil
}

func (d Device) getInfoUint(name string, id C.cl_device_info) (uint, error) {
	var buf C.cl_uint
	if err := C.clGetDeviceInfo(d.id, id, C.size_t(unsafe.Sizeof(buf)), unsafe.Pointer(&buf), nil); err != C.CL_SUCCESS {
		return 0, fmt.Errorf("error getting device info %s: %d", name, err)
	}

	return uint(buf), nil
}

func (d Device) getInfoUlong(name string, id C.cl_device_info) (uint64, error) {
	var buf C.cl_long
	if err := C.clGetDeviceInfo(d.id, id, C.size_t(unsafe.Sizeof(buf)), unsafe.Pointer(&buf), nil); err != C.CL_SUCCESS {
		return 0, fmt.Errorf("error getting device info %s: %d", name, err)
	}

	return uint64(buf), nil
}

func (d Device) getInfoSize(name string, id C.cl_device_info) (uintptr, error) {
	var buf C.size_t
	if err := C.clGetDeviceInfo(d.id, id, C.size_t(unsafe.Sizeof(buf)), unsafe.Pointer(&buf), nil); err != C.CL_SUCCESS {
		return 0, fmt.Errorf("error getting device info %s: %d", name, err)
	}

	return uintptr(buf), nil
}

func (d Device) getInfoSizeArray(name string, id C.cl_device_info) ([]uintptr, error) {
	var bytes C.size_t
	if err := C.clGetDeviceInfo(d.id, id, 0, nil, &bytes); err != C.CL_SUCCESS {
		return nil, fmt.Errorf("error getting length of device info %s: %d", name, err)
	}

	n := bytes / C.size_t(unsafe.Sizeof(C.size_t(0)))
	bytes = n * C.size_t(unsafe.Sizeof(C.size_t(0)))

	buf := make([]C.size_t, n)
	if err := C.clGetDeviceInfo(d.id, id, bytes, unsafe.Pointer(&buf[0]), nil); err != C.CL_SUCCESS {
		return nil, fmt.Errorf("error getting device info %s: %d", name, err)
	}

	result := make([]uintptr, n)
	for i, v := range buf {
		result[i] = uintptr(v)
	}

	return result, nil
}

func (d Device) getInfoString(name string, id C.cl_device_info) (string, error) {
	var n C.size_t
	if err := C.clGetDeviceInfo(d.id, id, 0, nil, &n); err != C.CL_SUCCESS {
		return "", fmt.Errorf("error getting length of device info %s: %d", name, err)
	}

	buf := make([]C.char, n)
	if err := C.clGetDeviceInfo(d.id, id, n, unsafe.Pointer(&buf[0]), nil); err != C.CL_SUCCESS {
		return "", fmt.Errorf("error getting device info %s: %d", name, err)
	}

	return C.GoString(&buf[0]), nil
}

func (d Device) getInfoFPConfig(name string, id C.cl_device_info) (FPConfig, error) {
	var buf C.cl_device_fp_config
	var result FPConfig
	if err := C.clGetDeviceInfo(d.id, id, C.size_t(unsafe.Sizeof(buf)), unsafe.Pointer(&buf), nil); err != C.CL_SUCCESS {
		return result, fmt.Errorf("error getting device info %s: %d", name, err)
	}

	result.config = buf
	return result, nil
}

func (d Device) getInfoMemCacheType(name string, id C.cl_device_info) (MemCacheType, error) {
	var buf C.cl_device_mem_cache_type
	if err := C.clGetDeviceInfo(d.id, id, C.size_t(unsafe.Sizeof(buf)), unsafe.Pointer(&buf), nil); err != C.CL_SUCCESS {
		return CacheNone, fmt.Errorf("error getting device inf %s: %d", name, err)
	}

	return memCacheType(buf)
}

func (d Device) getInfoLocalMemType(name string, id C.cl_device_info) (LocalMemType, error) {
	var buf C.cl_device_local_mem_type
	if err := C.clGetDeviceInfo(d.id, id, C.size_t(unsafe.Sizeof(buf)), unsafe.Pointer(&buf), nil); err != C.CL_SUCCESS {
		return LocalMemNone, fmt.Errorf("error getting device inf %s: %d", name, err)
	}

	return localMemType(buf)
}

func (d Device) getInfoQueueProperties(name string, id C.cl_device_info) (CommandQueueProperties, error) {
	var buf C.cl_command_queue_properties
	if err := C.clGetDeviceInfo(d.id, id, C.size_t(unsafe.Sizeof(buf)), unsafe.Pointer(&buf), nil); err != C.CL_SUCCESS {
		return CommandQueueProperties{}, fmt.Errorf("error getting device inf %s: %d", name, err)
	}

	return fromCLProperties(buf), nil
}

func (d Device) dummyUseStrings() []string {
	// a dummy function so that strings is not unused even if template space-delim is not used
	return strings.Split("", "")
}

func (d Device) Type() (DeviceType, error) {
	return d.getInfoType("Type", C.CL_DEVICE_TYPE)
}

func (d Device) VendorID() (uint, error) {
	return d.getInfoUint("VendorID", C.CL_DEVICE_VENDOR_ID)
}

func (d Device) MaxComputeUnits() (uint, error) {
	return d.getInfoUint("MaxComputeUnits", C.CL_DEVICE_MAX_COMPUTE_UNITS)
}

func (d Device) MaxWorkItemDimensions() (uint, error) {
	return d.getInfoUint("MaxWorkItemDimensions", C.CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS)
}

func (d Device) MaxWorkItemSizes() ([]uintptr, error) {
	return d.getInfoSizeArray("MaxWorkItemSizes", C.CL_DEVICE_MAX_WORK_ITEM_SIZES)
}

func (d Device) MaxWorkGroupSize() (uintptr, error) {
	return d.getInfoSize("MaxWorkGroupSize", C.CL_DEVICE_MAX_WORK_GROUP_SIZE)
}

func (d Device) PreferredVectorWidthChar() (uint, error) {
	return d.getInfoUint("PreferredVectorWidthChar", C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR)
}

func (d Device) PreferredVectorWidthShort() (uint, error) {
	return d.getInfoUint("PreferredVectorWidthShort", C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT)
}

func (d Device) PreferredVectorWidthInt() (uint, error) {
	return d.getInfoUint("PreferredVectorWidthInt", C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT)
}

func (d Device) PreferredVectorWidthLong() (uint, error) {
	return d.getInfoUint("PreferredVectorWidthLong", C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG)
}

func (d Device) PreferredVectorWidthFloat() (uint, error) {
	return d.getInfoUint("PreferredVectorWidthFloat", C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT)
}

func (d Device) PreferredVectorWidthDouble() (uint, error) {
	return d.getInfoUint("PreferredVectorWidthDouble", C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE)
}

func (d Device) PreferredVectorWidthHalf() (uint, error) {
	return d.getInfoUint("PreferredVectorWidthHalf", C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF)
}

func (d Device) NativeVectorWidthChar() (uint, error) {
	return d.getInfoUint("NativeVectorWidthChar", C.CL_DEVICE_NATIVE_VECTOR_WIDTH_CHAR)
}

func (d Device) NativeVectorWidthShort() (uint, error) {
	return d.getInfoUint("NativeVectorWidthShort", C.CL_DEVICE_NATIVE_VECTOR_WIDTH_SHORT)
}

func (d Device) NativeVectorWidthInt() (uint, error) {
	return d.getInfoUint("NativeVectorWidthInt", C.CL_DEVICE_NATIVE_VECTOR_WIDTH_INT)
}

func (d Device) NativeVectorWidthLong() (uint, error) {
	return d.getInfoUint("NativeVectorWidthLong", C.CL_DEVICE_NATIVE_VECTOR_WIDTH_LONG)
}

func (d Device) NativeVectorWidthFloat() (uint, error) {
	return d.getInfoUint("NativeVectorWidthFloat", C.CL_DEVICE_NATIVE_VECTOR_WIDTH_FLOAT)
}

func (d Device) NativeVectorWidthDouble() (uint, error) {
	return d.getInfoUint("NativeVectorWidthDouble", C.CL_DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE)
}

func (d Device) NativeVectorWidthHalf() (uint, error) {
	return d.getInfoUint("NativeVectorWidthHalf", C.CL_DEVICE_NATIVE_VECTOR_WIDTH_HALF)
}

func (d Device) MaxClockFrequency() (uint, error) {
	return d.getInfoUint("MaxClockFrequency", C.CL_DEVICE_MAX_CLOCK_FREQUENCY)
}

func (d Device) MaxMemAllocSize() (uint64, error) {
	return d.getInfoUlong("MaxMemAllocSize", C.CL_DEVICE_MAX_MEM_ALLOC_SIZE)
}

func (d Device) ImageSupport() (bool, error) {
	return d.getInfoBool("ImageSupport", C.CL_DEVICE_IMAGE_SUPPORT)
}

func (d Device) MaxReadImageArgs() (uint, error) {
	return d.getInfoUint("MaxReadImageArgs", C.CL_DEVICE_MAX_READ_IMAGE_ARGS)
}

func (d Device) MaxWriteImageArgs() (uint, error) {
	return d.getInfoUint("MaxWriteImageArgs", C.CL_DEVICE_MAX_WRITE_IMAGE_ARGS)
}

func (d Device) Image2DMaxWidth() (uintptr, error) {
	return d.getInfoSize("Image2DMaxWidth", C.CL_DEVICE_IMAGE2D_MAX_WIDTH)
}

func (d Device) Image2DMaxHeight() (uintptr, error) {
	return d.getInfoSize("Image2DMaxHeight", C.CL_DEVICE_IMAGE2D_MAX_HEIGHT)
}

func (d Device) Image3DMaxWidth() (uintptr, error) {
	return d.getInfoSize("Image3DMaxWidth", C.CL_DEVICE_IMAGE3D_MAX_WIDTH)
}

func (d Device) Image3DMaxHeight() (uintptr, error) {
	return d.getInfoSize("Image3DMaxHeight", C.CL_DEVICE_IMAGE3D_MAX_HEIGHT)
}

func (d Device) Image3DMaxDepth() (uintptr, error) {
	return d.getInfoSize("Image3DMaxDepth", C.CL_DEVICE_IMAGE3D_MAX_DEPTH)
}

func (d Device) ImageMaxBufferSize() (uintptr, error) {
	return d.getInfoSize("ImageMaxBufferSize", C.CL_DEVICE_IMAGE_MAX_BUFFER_SIZE)
}

func (d Device) ImageMaxArraySize() (uintptr, error) {
	return d.getInfoSize("ImageMaxArraySize", C.CL_DEVICE_IMAGE_MAX_ARRAY_SIZE)
}

func (d Device) MaxSamplers() (uint, error) {
	return d.getInfoUint("MaxSamplers", C.CL_DEVICE_MAX_SAMPLERS)
}

func (d Device) MaxParameterSize() (uintptr, error) {
	return d.getInfoSize("MaxParameterSize", C.CL_DEVICE_MAX_PARAMETER_SIZE)
}

func (d Device) MemBaseAddrAlign() (uint, error) {
	return d.getInfoUint("MemBaseAddrAlign", C.CL_DEVICE_MEM_BASE_ADDR_ALIGN)
}

func (d Device) SingleFPConfig() (FPConfig, error) {
	return d.getInfoFPConfig("SingleFPConfig", C.CL_DEVICE_SINGLE_FP_CONFIG)
}

func (d Device) DoubleFPConfig() (FPConfig, error) {
	return d.getInfoFPConfig("DoubleFPConfig", C.CL_DEVICE_SINGLE_FP_CONFIG)
}

func (d Device) GlobalMemCacheType() (MemCacheType, error) {
	return d.getInfoMemCacheType("GlobalMemCacheType", C.CL_DEVICE_GLOBAL_MEM_CACHE_TYPE)
}

func (d Device) GlobalMemCachelineSize() (uint, error) {
	return d.getInfoUint("GlobalMemCachelineSize", C.CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE)
}

func (d Device) GlobalMemCacheSize() (uint64, error) {
	return d.getInfoUlong("GlobalMemCacheSize", C.CL_DEVICE_GLOBAL_MEM_CACHE_SIZE)
}

func (d Device) GlobalMemSize() (uint64, error) {
	return d.getInfoUlong("GlobalMemSize", C.CL_DEVICE_GLOBAL_MEM_SIZE)
}

func (d Device) MaxConstantBufferSize() (uint64, error) {
	return d.getInfoUlong("MaxConstantBufferSize", C.CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE)
}

func (d Device) MaxConstantArgs() (uint, error) {
	return d.getInfoUint("MaxConstantArgs", C.CL_DEVICE_MAX_CONSTANT_ARGS)
}

func (d Device) LocalMemType() (LocalMemType, error) {
	return d.getInfoLocalMemType("LocalMemType", C.CL_DEVICE_LOCAL_MEM_TYPE)
}

func (d Device) LocalMemSize() (uint64, error) {
	return d.getInfoUlong("LocalMemSize", C.CL_DEVICE_LOCAL_MEM_SIZE)
}

func (d Device) ErrorCorrectionSupport() (bool, error) {
	return d.getInfoBool("ErrorCorrectionSupport", C.CL_DEVICE_ERROR_CORRECTION_SUPPORT)
}

func (d Device) ProfilingTimerResolution() (uintptr, error) {
	return d.getInfoSize("ProfilingTimerResolution", C.CL_DEVICE_PROFILING_TIMER_RESOLUTION)
}

func (d Device) HostUnifiedMemory() (bool, error) {
	return d.getInfoBool("HostUnifiedMemory", C.CL_DEVICE_HOST_UNIFIED_MEMORY)
}

func (d Device) EndianLittle() (bool, error) {
	return d.getInfoBool("EndianLittle", C.CL_DEVICE_ENDIAN_LITTLE)
}

func (d Device) Available() (bool, error) {
	return d.getInfoBool("Available", C.CL_DEVICE_AVAILABLE)
}

func (d Device) CompilerAvailable() (bool, error) {
	return d.getInfoBool("CompilerAvailable", C.CL_DEVICE_COMPILER_AVAILABLE)
}

func (d Device) LinkerAvailable() (bool, error) {
	return d.getInfoBool("LinkerAvailable", C.CL_DEVICE_LINKER_AVAILABLE)
}

func (d Device) QueueProperties() (CommandQueueProperties, error) {
	return d.getInfoQueueProperties("QueueProperties", C.CL_DEVICE_QUEUE_PROPERTIES)
}

func (d Device) BuiltInKernels() ([]string, error) {
	str, err := d.getInfoString("BuiltInKernels", C.CL_DEVICE_BUILT_IN_KERNELS)
	if err != nil {
		return nil, err
	}
	return strings.Split(str, ";"), nil
}

func (d Device) Name() (string, error) {
	return d.getInfoString("Name", C.CL_DEVICE_NAME)
}

func (d Device) Vendor() (string, error) {
	return d.getInfoString("Vendor", C.CL_DEVICE_VENDOR)
}

func (d Device) DriverVersion() (string, error) {
	return d.getInfoString("DriverVersion", C.CL_DRIVER_VERSION)
}

func (d Device) Profile() (string, error) {
	return d.getInfoString("Profile", C.CL_DEVICE_PROFILE)
}

func (d Device) Version() (string, error) {
	return d.getInfoString("Version", C.CL_DEVICE_VERSION)
}
