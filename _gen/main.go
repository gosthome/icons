package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"_gen/fontawesome"
	googlemdi "_gen/googleMDI"
	templarianmaterialdesignsvg "_gen/templarianMaterialDesignSVG"

	"github.com/Southclaws/fault/fctx"
	"golang.org/x/sync/errgroup"
)

func maybeRecover() {
	if err := recover(); err != nil {
		slog.Error("recoved from panic", "err", err)
	}
}

func wctx(ctx context.Context, f func(context.Context) error) func() error {
	return func() error {
		return f(ctx)
	}
}

func main() {
	err := os.Chdir("..")
	if err != nil {
		panic(err)
	}
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(wctx(ctx, googlemdi.GenerateIcons))
	eg.Go(wctx(ctx, templarianmaterialdesignsvg.GenerateIcons))
	eg.Go(wctx(ctx, fontawesome.GenerateIcons))
	err = eg.Wait()
	if err != nil {
		fmt.Print("context:\n")
		for k, v := range fctx.Unwrap(err) {
			fmt.Printf("\t%s:\t%s\n", k, v)
		}
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}
