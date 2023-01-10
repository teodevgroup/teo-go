package teo

import "fmt"

type AppBuilder struct {
	Name string
}

func NewAppBuilder(name string) *AppBuilder {
	return &AppBuilder{name}
}

func makeAppBuilder(name string) AppBuilder {
	return AppBuilder{name}
}

func (appBuilder AppBuilder) Load() {
	return
}

func (appBuilder AppBuilder) Build() {
	return
}

func Hello() {
	fmt.Println("Hello, World!")
}
