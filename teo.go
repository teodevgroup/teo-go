package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L../teo-ffi/target/release -lteo_ffi
#include "teo.h"
*/
import "C"
import (
	"runtime"
	"sync"
)

type AppBuilder struct {
	cAppBuilder C.struct__AppBuilder
}

func newCAppBuilder() C.struct__AppBuilder {
	return (C.struct__AppBuilder)(C.AppBuilder_Create(C.CString("Go"), C.CString(runtime.Version()[2:])))
}

func NewApp() *AppBuilder {
	return &AppBuilder{
		cAppBuilder: newCAppBuilder(),
	}
}

func (appBuilder AppBuilder) Run() {
	var wg sync.WaitGroup
	wg.Add(1)
	C.AppBuilder_Build(&appBuilder.cAppBuilder)
	wg.Wait()
}

func main() {
	app := NewApp()
	app.Run()
}
