//go:build cuda

// CUDA cgo linkage for Molequla, split out of cgo_aml.go so the default
// build is CPU-only. `go build` → CPU path (OpenBLAS). `go build -tags cuda`
// → adds the CUDA backend (notorch_cuda.o + cuBLAS); requires an
// nvcc-built notorch_cuda.o present and a CUDA toolchain on the host.
// Paired with the `linux && cuda` / `!linux || !cuda` build tags on
// gpu_forward{,_stub}.go and gpu_bindings_{linux,stub}.go.

package main

/*
#cgo linux CFLAGS: -DUSE_CUDA -I/usr/local/cuda/include
#cgo linux LDFLAGS: ${SRCDIR}/ariannamethod/notorch_cuda.o -L/usr/local/cuda/lib64 -lcudart -lcublas -lstdc++
*/
import "C"
