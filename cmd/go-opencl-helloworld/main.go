package main

import (
	"fmt"
	"log"

	"github.com/Patagonicus/opencl"
)

const source string = `__kernel void hello(__constant char *a, __constant char *b, __global char *c) {
	size_t gid = get_global_id(0);
	if (gid < get_global_size(0)) {
		c[gid] = a[gid] + b[gid];
	}
}`
const kernelName = "hello"

var dataA = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8}
var dataB = []byte{0, 1, 2, 3, 4, 3, 2, 1, 0}

func getFirstDevice() (*opencl.Device, error) {
	platforms, err := opencl.GetPlatforms()
	if err != nil {
		return nil, err
	}

	for _, p := range platforms {
		devices, err := p.Devices(opencl.DeviceTypeAll)
		if err != nil {
			return nil, err
		}

		if len(devices) > 0 {
			return &devices[0], nil
		}
	}

	return nil, fmt.Errorf("no device found")
}

func getDeviceContext() (*opencl.Device, *opencl.Context, error) {
	device, err := getFirstDevice()
	if err != nil {
		return nil, nil, err
	}

	context, err := opencl.CreateContext(nil, []*opencl.Device{device}, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	return device, context, nil
}

func createBuffers(context opencl.Context) (*opencl.Memory, *opencl.Memory, *opencl.Memory, error) {
	bufA, err := context.CreateBuffer(opencl.MemReadOnly, uintptr(len(dataA)), nil)
	if err != nil {
		return nil, nil, nil, err
	}

	bufB, err := context.CreateBuffer(opencl.MemReadOnly, uintptr(len(dataA)), nil)
	if err != nil {
		bufA.Release()
		return nil, nil, nil, err
	}

	bufC, err := context.CreateBuffer(opencl.MemWriteOnly, uintptr(len(dataA)), nil)
	if err != nil {
		bufA.Release()
		bufB.Release()
		return nil, nil, nil, err
	}

	return bufA, bufB, bufC, nil
}

func createKernel(context opencl.Context, device opencl.Device) (*opencl.Program, *opencl.Kernel, error) {
	program, err := context.CreateProgramWithSource([]string{source})
	if err != nil {
		return nil, nil, err
	}

	err = program.Build([]*opencl.Device{&device}, "")
	if err != nil {
		buildLog, logErr := program.BuildLog(device)
		if logErr == nil {
			log.Printf("build log: %s", buildLog)
		}
		program.Release()
		return nil, nil, err
	}

	kernel, err := program.CreateKernel(kernelName)
	if err != nil {
		program.Release()
		return nil, nil, err
	}

	return program, kernel, nil
}

func main() {
	log.SetFlags(0)

	device, context, err := getDeviceContext()
	if err != nil {
		log.Fatal(err)
	}
	defer context.Release()

	bufA, bufB, bufC, err := createBuffers(*context)
	if err != nil {
		log.Fatal(err)
	}
	defer bufA.Release()
	defer bufB.Release()
	defer bufC.Release()

	queue, err := context.CreateCommandQueue(*device, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer queue.Release()

	err = queue.EnqueueWriteBuffer(*bufA, 0, dataA, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = queue.EnqueueWriteBuffer(*bufB, 0, dataB, nil)
	if err != nil {
		log.Fatal(err)
	}

	program, kernel, err := createKernel(*context, *device)
	if err != nil {
		log.Fatal(err)
	}
	defer program.Release()
	defer kernel.Release()

	err = kernel.SetArgBuffer(0, *bufA)
	if err != nil {
		log.Fatal(err)
	}
	err = kernel.SetArgBuffer(1, *bufB)
	if err != nil {
		log.Fatal(err)
	}
	err = kernel.SetArgBuffer(2, *bufC)
	if err != nil {
		log.Fatal(err)
	}

	event, err := queue.EnqueueNDRangeKernel(*kernel, []uintptr{0}, []uintptr{uintptr(len(dataA))}, []uintptr{uintptr(len(dataA))}, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer event.Release()

	data := make([]byte, len(dataA))
	err = queue.EnqueueReadBuffer(*bufC, 0, data, []opencl.Event{*event})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%d\n", data)
}
