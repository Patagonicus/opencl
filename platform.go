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

type Platform struct {
	id C.cl_platform_id
}

func GetPlatforms() ([]Platform, error) {
	var num_platforms C.cl_uint
	if err := C.clGetPlatformIDs(0, nil, &num_platforms); err != C.CL_SUCCESS {
		return nil, fmt.Errorf("error getting number of platforms: %d", err)
	}

	ids := make([]C.cl_platform_id, num_platforms)
	if err := C.clGetPlatformIDs(num_platforms, &ids[0], nil); err != C.CL_SUCCESS {
		return nil, fmt.Errorf("error getting platforms: %d", err)
	}

	platforms := make([]Platform, num_platforms)
	for i, id := range ids {
		platforms[i].id = id
	}

	return platforms, nil
}

func (p Platform) Devices(deviceType DeviceType) ([]Device, error) {
	return getDevices(p, deviceType)
}
