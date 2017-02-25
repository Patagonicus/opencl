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

var templates = []struct {
	in   string
	out  string
	data interface{}
}{
	{"platforminfo.go.tmpl", "platforminfo.go", platformInfo},
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

	return buf.Bytes(), nil
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
