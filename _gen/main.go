package main

import (
	"log/slog"
	"os"
	"sync"

	"_gen/fontawesome"
	googlemdi "_gen/googleMDI"
	templarianmaterialdesignsvg "_gen/templarianMaterialDesignSVG"
)

func maybeRecover() {
	if err := recover(); err != nil {
		slog.Error("recoved from panic", "err", err)
	}
}

func main() {
	err := os.Chdir("..")
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		googlemdi.GenerateIcons()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()

		templarianmaterialdesignsvg.GenerateIcons()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()

		fontawesome.GenerateIcons()
	}()
	wg.Wait()
}
