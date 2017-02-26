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
	"sort"
	"strings"
)

var fpconfigNames = map[C.cl_device_fp_config]string{
	C.CL_FP_DENORM:           "denorm",
	C.CL_FP_INF_NAN:          "inf NaN",
	C.CL_FP_ROUND_TO_NEAREST: "round to nearest",
	C.CL_FP_ROUND_TO_ZERO:    "round to zero",
	C.CL_FP_ROUND_TO_INF:     "round to inf",
	C.CL_FP_FMA:              "fused multiply-add",
	C.CL_FP_SOFT_FLOAT:       "software",
}

type FPConfig struct {
	config C.cl_device_fp_config
}

func (c FPConfig) Denorm() bool {
	return (c.config & C.CL_FP_DENORM) != 0
}

func (c FPConfig) InfNaN() bool {
	return (c.config & C.CL_FP_INF_NAN) != 0
}

func (c FPConfig) RoundToNearest() bool {
	return (c.config & C.CL_FP_ROUND_TO_NEAREST) != 0
}

func (c FPConfig) RoundToZero() bool {
	return (c.config & C.CL_FP_ROUND_TO_ZERO) != 0
}

func (c FPConfig) RoundToInf() bool {
	return (c.config & C.CL_FP_ROUND_TO_INF) != 0
}

func (c FPConfig) FMA() bool {
	return (c.config & C.CL_FP_FMA) != 0
}

func (c FPConfig) CorrectlyRoundedDivideSqrt() bool {
	return (c.config & C.CL_FP_CORRECTLY_ROUNDED_DIVIDE_SQRT) != 0
}

func (c FPConfig) SoftFloat() bool {
	return (c.config & C.CL_FP_SOFT_FLOAT) != 0
}

func (c FPConfig) String() string {
	var features []string
	for bitmask, name := range fpconfigNames {
		if (c.config & bitmask) != 0 {
			features = append(features, name)
		}
	}

	sort.Strings(features)

	return fmt.Sprintf("FPConfig{%s}", strings.Join(features, ", "))
}
