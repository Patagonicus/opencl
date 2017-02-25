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

type Device struct {
	id C.cl_device_id
}

type DeviceType C.cl_device_type

const (
	DeviceTypeCPU         = DeviceType(C.CL_DEVICE_TYPE_CPU)
	DeviceTypeGPU         = DeviceType(C.CL_DEVICE_TYPE_GPU)
	DeviceTypeAccelerator = DeviceType(C.CL_DEVICE_TYPE_ACCELERATOR)
	DeviceTypeCustom      = DeviceType(C.CL_DEVICE_TYPE_CUSTOM)
	DeviceTypeDefault     = DeviceType(C.CL_DEVICE_TYPE_DEFAULT)
	DeviceTypeAll         = DeviceType(C.CL_DEVICE_TYPE_ALL)
)

func (t DeviceType) String() string {
	switch t {
	case DeviceTypeCPU:
		return "CPU"
	case DeviceTypeGPU:
		return "GPU"
	case DeviceTypeAccelerator:
		return "accelereator"
	case DeviceTypeCustom:
		return "custom"
	case DeviceTypeDefault:
		return "default"
	case DeviceTypeAll:
		return "all"
	}
	return fmt.Sprintf("unknown type %d", t)
}

func getDevices(platform Platform, deviceType DeviceType) ([]Device, error) {
	var n C.cl_uint
	if err := C.clGetDeviceIDs(platform.id, C.cl_device_type(deviceType), 0, nil, &n); err != C.CL_SUCCESS {
		return nil, fmt.Errorf("error getting number of devices: %d", err)
	}

	ids := make([]C.cl_device_id, n)
	if err := C.clGetDeviceIDs(platform.id, C.cl_device_type(deviceType), n, unsafe.Pointer(&ids[0]), nil); err != C.CL_SUCCESS {
		return nil, fmt.Errorf("error getting devices: %d", err)
	}

	devices := make([]Device, n)
	for i, id := range ids {
		devices[i].id = id
	}

	return devices, nil
}
