# go-opencl

go-opencl is a wrapper around the C OpenCL 1.2 API for Go. It aims to be close to the C API while adding Go features (object orientation, `[]string` instead of space delimited extension list, etc.).

It is mainly developed and tested on Linux and with Intel and Nvidia GPUs. It should work with AMD GPUs and, at least in theory, it should work on Windows and macOS.

## Features

- [x] querying platforms
  - [x] enumerate platforms
  - [x] query info
- [x] querying devices
  - [x] enumerate devices of a platform
  - [x] querying info
- [ ] subdevices (not planned for now)
- [ ] context
  - [x] creating and releasing contexts
  - [ ] querying info
- [ ] command queues
  - [x] creating and releasing command queues
  - [ ] querying info
- [ ] buffers
  - [*] creating and releasing buffers
  - [x] writing buffers (blocking only for now)
  - [x] reading buffers (blocking only for now)
  - [ ] querying info
- [ ] program objects
  - [x] creating and releasing program objects
  - [x] building program executables
  - [ ] separate compiling and linking (not planned for now)
  - [ ] unloading the compiler
  - [ ] querying program (build) info
- [ ] kernel objects
  - [x] creating and releasing kernels
  - [x] setting arguments
  - [ ] querying kernel, work group and argument info
  - [x] executing kernels
- [ ] event objects
  - [ ] creating and releasing user events (not planned for now)
  - [ ] query event info
  - [ ] flush and finish
  - [ ] wait for events
  - [ ] event callbacks (not planned for now)
  - [ ] markers and barriers (not planned for now)
  - [ ] profiling info (not planned for now)
- [ ] images (not planned for now)
  - [ ] write detailed todo for images
- [ ] Direct3D sharing (not planned for now)
- [ ] OpenGL sharing (not planned for now)
- [ ] DX9 Media Surface sharing (not planned for now)
- [ ] Direct3D 11 sharing (not planned for now)

## Documentation

There isn't too much documentation in this project, however the API mimics the C API (with some object orientation thrown in), so by taking a look at the official OpenCL 1.2 documentation it should be easy enough to understand. Take a look at the included `go-opencl-info` utility to see how to use this API.
