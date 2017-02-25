package opencl

//go:generate go run infogen.go

// #cgo LDFLAGS: -lOpenCL
/*
#ifdef __APPLE__
#include <OpenCL/opencl.h>
#else
#include <CL/cl.h>
#endif
*/
import "C"
