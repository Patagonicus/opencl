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

type CommandQueue struct {
	queue *C.cl_command_queue
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
	var queue CommandQueue
	queue.queue = (*C.cl_command_queue)(C.malloc(C.sizeof_cl_command_queue))
	var err C.cl_int
	*queue.queue = C.clCreateCommandQueue(*context.context, device.id, clProperties, &err)
	if err != C.CL_SUCCESS {
		C.free(unsafe.Pointer(queue.queue))
		return nil, fmt.Errorf("failed to create command queue: %d", err)
	}
	return &queue, nil
}

func (q CommandQueue) Release() error {
	err := C.clReleaseCommandQueue(*q.queue)
	if err != C.CL_SUCCESS {
		return fmt.Errorf("falied to release command queue: %d", err)
	}
	C.free(unsafe.Pointer(q.queue))
	return nil
}

func (q CommandQueue) EnqueueReadBuffer(memory Memory, offset uintptr, buffer []byte, waitList []Event) error {
	var clEventsPtr unsafe.Pointer
	if len(waitList) > 0 {
		clEventsPtr = unsafe.Pointer(&asCLEventList(waitList)[0])
	}
	if err := C.clEnqueueReadBuffer(*q.queue, *memory.memory, C.CL_TRUE, C.size_t(offset), C.size_t(len(buffer)), unsafe.Pointer(&buffer[0]), C.cl_uint(len(waitList)), clEventsPtr, nil); err != C.CL_SUCCESS {
		return fmt.Errorf("failed to enqueue read: %d", err)
	}
	return nil
}

func (q CommandQueue) EnqueueWriteBuffer(memory Memory, offset uintptr, buffer []byte, waitList []Event) error {
	var clEventsPtr unsafe.Pointer
	if len(waitList) > 0 {
		clEventsPtr = unsafe.Pointer(&asCLEventList(waitList)[0])
	}
	if err := C.clEnqueueWriteBuffer(*q.queue, *memory.memory, C.CL_TRUE, C.size_t(offset), C.size_t(len(buffer)), unsafe.Pointer(&buffer[0]), C.cl_uint(len(waitList)), clEventsPtr, nil); err != C.CL_SUCCESS {
		return fmt.Errorf("Failed to enqueue write: %d", err)
	}
	return nil
}

func (q CommandQueue) EnqueueNDRangeKernel(kernel Kernel, globalWorkOffset, globalWorkSize, localWorkSize []uintptr, waitList []Event) (*Event, error) {
	if len(globalWorkOffset) != len(globalWorkSize) || len(globalWorkOffset) != len(localWorkSize) {
		return nil, fmt.Errorf("globalWorkOffset, globalWorkSize and localWorkSize have to have the same length")
	}

	// TODO: check that len(dim) > 0
	dim := C.cl_uint(len(globalWorkOffset))
	clGlobalWorkOffset := make([]C.size_t, dim)
	clGlobalWorkSize := make([]C.size_t, dim)
	clLocalWorkSize := make([]C.size_t, dim)
	for i := range globalWorkOffset {
		clGlobalWorkOffset[i] = C.size_t(globalWorkOffset[i])
		clGlobalWorkSize[i] = C.size_t(globalWorkSize[i])
		clLocalWorkSize[i] = C.size_t(localWorkSize[i])
	}

	var clEventsPtr *C.cl_event
	if len(waitList) > 0 {
		clEventsPtr = &asCLEventList(waitList)[0]
	}

	var event Event
	err := C.clEnqueueNDRangeKernel(*q.queue, *kernel.kernel, dim, &clGlobalWorkOffset[0], &clGlobalWorkSize[0], &clLocalWorkSize[0], C.cl_uint(len(waitList)), clEventsPtr, &event.event)
	if err != C.CL_SUCCESS {
		return nil, fmt.Errorf("failed to enqueue kernel: %d", err)
	}

	return &event, nil
}
