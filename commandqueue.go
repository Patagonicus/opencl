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

type CommandQueue struct {
	queue C.cl_command_queue
}

type CommandQueueProperties struct {
	OutOfOrderExec bool
	Profiling      bool
}

func fromCLProperties(properties C.cl_command_queue_properties) CommandQueueProperties {
	var result CommandQueueProperties
	result.OutOfOrderExec = ((properties & C.CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE) != 0)
	result.Profiling = ((properties & C.CL_QUEUE_PROFILING_ENABLE) != 0)
	return result
}

func (p CommandQueueProperties) toCLProperties() C.cl_command_queue_properties {
	var result C.cl_command_queue_properties
	if p.OutOfOrderExec {
		result |= C.CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE
	}
	if p.Profiling {
		result |= C.CL_QUEUE_PROFILING_ENABLE
	}
	return result
}

func (p CommandQueueProperties) String() string {
	return fmt.Sprintf("QueueProperties{OutOfOrderExec: %t, Profiling: %t}", p.OutOfOrderExec, p.Profiling)
}

func createCommandQueue(context Context, device Device, properties *CommandQueueProperties) (*CommandQueue, error) {
	var clProperties C.cl_command_queue_properties
	if properties != nil {
		clProperties = properties.toCLProperties()
	}
	var err C.cl_int
	queue := C.clCreateCommandQueue(context.context, device.id, clProperties, &err)
	if err != C.CL_SUCCESS {
		return nil, fmt.Errorf("failed to create command queue: %d", err)
	}
	return &CommandQueue{queue}, nil
}

func (q CommandQueue) Release() error {
	err := C.clReleaseCommandQueue(q.queue)
	if err != C.CL_SUCCESS {
		return fmt.Errorf("falied to release command queue: %d", err)
	}
	return nil
}
