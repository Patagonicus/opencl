package opencl

/*
#ifdef __APPLE__
#include <OpenCL/opencl.h>
#else
#include <CL/cl.h>
#endif
*/
import "C"
import "fmt"

type Event struct {
	event C.cl_event
}

func (e Event) Release() error {
	err := C.clReleaseEvent(e.event)
	if err != C.CL_SUCCESS {
		return fmt.Errorf("failed to release event: %d", err)
	}
	return nil
}

func asCLEventList(events []Event) []C.cl_event {
	result := make([]C.cl_event, len(events))
	for i, e := range events {
		result[i] = e.event
	}
	return result
}
