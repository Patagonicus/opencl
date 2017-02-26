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

func (p CommandQueueProperties) String() string {
	return fmt.Sprintf("QueueProperties{OutOfOrderExec: %t, Profiling: %t}", p.OutOfOrderExec, p.Profiling)
}
