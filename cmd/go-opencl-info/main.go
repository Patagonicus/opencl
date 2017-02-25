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

		for i := range devices {
			log.Printf("  Device %d", i)
		}
	}
}
