package main

import (
	"log"
	"strings"

	"github.com/Patagonicus/opencl"
)

func main() {
	platforms, err := opencl.GetPlatforms()
	if err != nil {
		log.Fatal(err)
	}

	for i, p := range platforms {
		log.Printf("Platform %d", i)

		var name, vendor, profile, version string
		var extensions []string
		var err error

		if name, err = p.Name(); err != nil {
			log.Fatal(err)
		}
		if vendor, err = p.Vendor(); err != nil {
			log.Fatal(err)
		}
		if profile, err = p.Profile(); err != nil {
			log.Fatal(err)
		}
		if version, err = p.Version(); err != nil {
			log.Fatal(err)
		}
		if extensions, err = p.Extension(); err != nil {
			log.Fatal(err)
		}

		log.Printf("  Name: %s", name)
		log.Printf("  Vendor: %s", vendor)
		log.Printf("  Profile: %s", profile)
		log.Printf("  Version: %s", version)
		log.Printf("  Extensions: %s", strings.Join(extensions, ", "))

		devices, err := p.Devices(opencl.DeviceTypeAll)
		if err != nil {
			log.Fatal(err)
		}

		for i, d := range devices {
			log.Printf("  Device %d", i)

			var deviceType opencl.DeviceType
			var vendorID uint
			var maxWorkItemSizes []uintptr
			var maxWorkGroupSize uintptr
			var imageSupport bool
			var singleFP, doubleFP opencl.FPConfig
			var memCacheType opencl.MemCacheType
			var localMemType opencl.LocalMemType
			var queueProperties opencl.CommandQueueProperties
			var name, version, vendor, profile, driverVersion string
			var builtInKernels []string
			var err error

			if name, err = d.Name(); err != nil {
				log.Fatal(err)
			}
			if version, err = d.Version(); err != nil {
				log.Fatal(err)
			}
			if deviceType, err = d.Type(); err != nil {
				log.Fatal(err)
			}
			if profile, err = d.Profile(); err != nil {
				log.Fatal(err)
			}
			if vendor, err = d.Vendor(); err != nil {
				log.Fatal(err)
			}
			if vendorID, err = d.VendorID(); err != nil {
				log.Fatal(err)
			}
			if driverVersion, err = d.DriverVersion(); err != nil {
				log.Fatal(err)
			}
			if maxWorkItemSizes, err = d.MaxWorkItemSizes(); err != nil {
				log.Fatal(err)
			}
			if maxWorkGroupSize, err = d.MaxWorkGroupSize(); err != nil {
				log.Fatal(err)
			}
			if imageSupport, err = d.ImageSupport(); err != nil {
				log.Fatal(err)
			}
			if singleFP, err = d.SingleFPConfig(); err != nil {
				log.Fatal(err)
			}
			if doubleFP, err = d.DoubleFPConfig(); err != nil {
				log.Fatal(err)
			}
			if memCacheType, err = d.GlobalMemCacheType(); err != nil {
				log.Fatal(err)
			}
			if localMemType, err = d.LocalMemType(); err != nil {
				log.Fatal(err)
			}
			if queueProperties, err = d.QueueProperties(); err != nil {
				log.Fatal(err)
			}
			if builtInKernels, err = d.BuiltInKernels(); err != nil {
				log.Fatal(err)
			}

			log.Printf("    Name: %s", name)
			log.Printf("    Version: %s", version)
			log.Printf("    Type: %s", deviceType)
			log.Printf("    Profile: %s", profile)
			log.Printf("    Vendor: %s (ID: %d)", vendor, vendorID)
			log.Printf("    Driver version: %s", driverVersion)
			log.Printf("    Max work item sizes: %d", maxWorkItemSizes)
			log.Printf("    Max work group size: %d", maxWorkGroupSize)
			log.Printf("    Image support: %t", imageSupport)
			log.Printf("    Single Floating Point: %s", singleFP)
			log.Printf("    Double Floating Point: %s", doubleFP)
			log.Printf("    Global memory cache type: %s", memCacheType.Name())
			log.Printf("    Local memory type: %s", localMemType.Name())
			log.Printf("    Queue properties: %s", queueProperties)
			log.Printf("    Built in kernels: %s", strings.Join(builtInKernels, " "))
		}
	}
}
