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

func (p Platform) getInfoString(name string, id C.cl_platform_info) (string, error) {
  var n C.size_t
  if err := C.clGetPlatformInfo(p.id, id, 0, nil, &n); err != C.CL_SUCCESS {
    return "", fmt.Errorf("error getting length of platform info %s: %d", name, err)
  }

  buf := make([]C.char, n)
  if err := C.clGetPlatformInfo(p.id, id, n, unsafe.Pointer(&buf[0]), nil); err != C.CL_SUCCESS {
    return "", fmt.Errorf("error getting platform info %s: %d", name, err)
  }

  return C.GoString(&buf[0]), nil
}

func (p Platform) dummyUseStrings() ([]string) {
  // a dummy function so that strings is not unused even if template space-delim is not used
  return strings.Split("", "")
}

func (p Platform) Profile() (string, error) {
  return p.getInfoString("Profile", C.CL_PLATFORM_PROFILE)
}

func (p Platform) Version() (string, error) {
  return p.getInfoString("Version", C.CL_PLATFORM_VERSION)
}

func (p Platform) Name() (string, error) {
  return p.getInfoString("Name", C.CL_PLATFORM_NAME)
}

func (p Platform) Vendor() (string, error) {
  return p.getInfoString("Vendor", C.CL_PLATFORM_VENDOR)
}

func (p Platform) Extension() ([]string, error) {
  str, err := p.getInfoString("Extension", C.CL_PLATFORM_EXTENSIONS)
  if err != nil {
    return nil, err
  }
  return strings.Split(str, " "), nil
}

