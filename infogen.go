//+build ignore

package main

/*
#ifdef __APPLE__
#include <OpenCL/opencl.h>
#else
#include <CL/cl.h>
#endif
*/
import "C"
import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

var platformInfo = []struct {
	Name       string
	ID         string
	ReturnType string
}{
	{"Profile", "C.CL_PLATFORM_PROFILE", "string"},
	{"Version", "C.CL_PLATFORM_VERSION", "string"},
	{"Name", "C.CL_PLATFORM_NAME", "string"},
	{"Vendor", "C.CL_PLATFORM_VENDOR", "string"},
	{"Extension", "C.CL_PLATFORM_EXTENSIONS", "space-delim"},
}

var deviceInfo = []struct {
	Name       string
	ID         string
	ReturnType string
}{
	{"Type", "C.CL_DEVICE_TYPE", "type"},
	{"VendorID", "C.CL_DEVICE_VENDOR_ID", "uint"},
	{"MaxComputeUnits", "C.CL_DEVICE_MAX_COMPUTE_UNITS", "uint"},
	{"MaxWorkItemDimensions", "C.CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS", "uint"},
	{"MaxWorkItemSizes", "C.CL_DEVICE_MAX_WORK_ITEM_SIZES", "size-array"},
	{"MaxWorkGroupSize", "C.CL_DEVICE_MAX_WORK_GROUP_SIZE", "size"},
	{"PreferredVectorWidthChar", "C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR", "uint"},
	{"PreferredVectorWidthShort", "C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT", "uint"},
	{"PreferredVectorWidthInt", "C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT", "uint"},
	{"PreferredVectorWidthLong", "C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG", "uint"},
	{"PreferredVectorWidthFloat", "C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT", "uint"},
	{"PreferredVectorWidthDouble", "C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE", "uint"},
	{"PreferredVectorWidthHalf", "C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF", "uint"},
	{"NativeVectorWidthChar", "C.CL_DEVICE_NATIVE_VECTOR_WIDTH_CHAR", "uint"},
	{"NativeVectorWidthShort", "C.CL_DEVICE_NATIVE_VECTOR_WIDTH_SHORT", "uint"},
	{"NativeVectorWidthInt", "C.CL_DEVICE_NATIVE_VECTOR_WIDTH_INT", "uint"},
	{"NativeVectorWidthLong", "C.CL_DEVICE_NATIVE_VECTOR_WIDTH_LONG", "uint"},
	{"NativeVectorWidthFloat", "C.CL_DEVICE_NATIVE_VECTOR_WIDTH_FLOAT", "uint"},
	{"NativeVectorWidthDouble", "C.CL_DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE", "uint"},
	{"NativeVectorWidthHalf", "C.CL_DEVICE_NATIVE_VECTOR_WIDTH_HALF", "uint"},
	{"MaxClockFrequency", "C.CL_DEVICE_MAX_CLOCK_FREQUENCY", "uint"},
	{"MaxMemAllocSize", "C.CL_DEVICE_MAX_MEM_ALLOC_SIZE", "ulong"},
	{"ImageSupport", "C.CL_DEVICE_IMAGE_SUPPORT", "bool"},
	{"MaxReadImageArgs", "C.CL_DEVICE_MAX_READ_IMAGE_ARGS", "uint"},
	{"MaxWriteImageArgs", "C.CL_DEVICE_MAX_WRITE_IMAGE_ARGS", "uint"},
	{"Image2DMaxWidth", "C.CL_DEVICE_IMAGE2D_MAX_WIDTH", "size"},
	{"Image2DMaxHeight", "C.CL_DEVICE_IMAGE2D_MAX_HEIGHT", "size"},
	{"Image3DMaxWidth", "C.CL_DEVICE_IMAGE3D_MAX_WIDTH", "size"},
	{"Image3DMaxHeight", "C.CL_DEVICE_IMAGE3D_MAX_HEIGHT", "size"},
	{"Image3DMaxDepth", "C.CL_DEVICE_IMAGE3D_MAX_DEPTH", "size"},
	{"ImageMaxBufferSize", "C.CL_DEVICE_IMAGE_MAX_BUFFER_SIZE", "size"},
	{"ImageMaxArraySize", "C.CL_DEVICE_IMAGE_MAX_ARRAY_SIZE", "size"},
	{"MaxSamplers", "C.CL_DEVICE_MAX_SAMPLERS", "uint"},
	{"MaxParameterSize", "C.CL_DEVICE_MAX_PARAMETER_SIZE", "size"},
	{"MemBaseAddrAlign", "C.CL_DEVICE_MEM_BASE_ADDR_ALIGN", "uint"},
	{"SingleFPConfig", "C.CL_DEVICE_SINGLE_FP_CONFIG", "fpconfig"},
	{"DoubleFPConfig", "C.CL_DEVICE_SINGLE_FP_CONFIG", "fpconfig"},
	{"GlobalMemCacheType", "C.CL_DEVICE_GLOBAL_MEM_CACHE_TYPE", "memcachetype"},
	{"GlobalMemCachelineSize", "C.CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE", "uint"},
	{"GlobalMemCacheSize", "C.CL_DEVICE_GLOBAL_MEM_CACHE_SIZE", "ulong"},
	{"GlobalMemSize", "C.CL_DEVICE_GLOBAL_MEM_SIZE", "ulong"},
	{"MaxConstantBufferSize", "C.CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE", "ulong"},
	{"MaxConstantArgs", "C.CL_DEVICE_MAX_CONSTANT_ARGS", "uint"},
	{"LocalMemType", "C.CL_DEVICE_LOCAL_MEM_TYPE", "localmemtype"},
	{"LocalMemSize", "C.CL_DEVICE_LOCAL_MEM_SIZE", "ulong"},
	{"ErrorCorrectionSupport", "C.CL_DEVICE_ERROR_CORRECTION_SUPPORT", "bool"},
	{"ProfilingTimerResolution", "C.CL_DEVICE_PROFILING_TIMER_RESOLUTION", "size"},
	{"HostUnifiedMemory", "C.CL_DEVICE_HOST_UNIFIED_MEMORY", "bool"},
	{"EndianLittle", "C.CL_DEVICE_ENDIAN_LITTLE", "bool"},
	{"Available", "C.CL_DEVICE_AVAILABLE", "bool"},
	{"CompilerAvailable", "C.CL_DEVICE_COMPILER_AVAILABLE", "bool"},
	{"LinkerAvailable", "C.CL_DEVICE_LINKER_AVAILABLE", "bool"},
	{"QueueProperties", "C.CL_DEVICE_QUEUE_PROPERTIES", "queueproperties"},
	{"BuiltInKernels", "C.CL_DEVICE_BUILT_IN_KERNELS", "semicolon-delim"},
	{"Name", "C.CL_DEVICE_NAME", "string"},
	{"Vendor", "C.CL_DEVICE_VENDOR", "string"},
	{"DriverVersion", "C.CL_DRIVER_VERSION", "string"},
	{"Profile", "C.CL_DEVICE_PROFILE", "string"},
	{"Version", "C.CL_DEVICE_VERSION", "string"},
}

var templates = []struct {
	in   string
	out  string
	data interface{}
}{
	{"platforminfo.go.tmpl", "platforminfo.go", platformInfo},
	{"deviceinfo.go.tmpl", "deviceinfo.go", deviceInfo},
}

func loadTemplate(filename string) (*template.Template, error) {
	return template.ParseFiles(filename)
}

func generate(tmpl *template.Template, data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}

	return format.Source(buf.Bytes())
}

func writeBytes(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

func generateInfo(templatename, filename string, data interface{}) error {
	tmpl, err := loadTemplate(templatename)
	if err != nil {
		return err
	}

	code, err := generate(tmpl, data)
	if err != nil {
		return err
	}

	return writeBytes(filename, code)
}

func main() {
	for _, t := range templates {
		err := generateInfo(t.in, t.out, t.data)
		if err != nil {
			os.Remove(t.out)
			log.Fatal(err)
		}
	}
}
