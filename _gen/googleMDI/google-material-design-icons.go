package googlemdi

import (
	"_gen/common"
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
)

func GenerateIcons(ctx context.Context) error {
	mdi_dir := "../_mdi.source"
	done := ctx.Done()
	ctx = fctx.WithMeta(ctx, "mdi_dir", mdi_dir)
	if _, err := os.Stat(mdi_dir); os.IsNotExist(err) {
		c := exec.CommandContext(ctx, "git", "clone", "--depth", "1", "https://github.com/google/material-design-icons.git", mdi_dir)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		select {
		case <-done:
			return fault.Wrap(ctx.Err())
		default:
		}
		err := c.Run()
		if err != nil {
			return fault.Wrap(err)
		}
		select {
		case <-done:
			return fault.Wrap(ctx.Err())
		default:
		}
	}
	files, err := filepath.Glob(filepath.Join(mdi_dir, "src/*/*/"))
	if err != nil {
		return fault.Wrap(err)
	}

	collections := map[string]*common.Collection{
		"materialicons": common.NewColl(
			mdi_dir,
			"google/materialdesignicons/icons.go",
		),
		"materialiconsoutlined": common.NewColl(
			mdi_dir,
			"google/materialdesigniconsoutlined/icons.go",
		),
		"materialiconsround": common.NewColl(
			mdi_dir,
			"google/materialdesigniconsround/icons.go",
		),
		"materialiconssharp": common.NewColl(
			mdi_dir,
			"google/materialdesigniconssharp/icons.go",
		),
	}
	for _, file := range files {
		base := filepath.Base(file)
		for collection, col := range collections {
			select {
			case <-done:
				return fault.Wrap(ctx.Err())
			default:
			}
			fp, err := filepath.Glob(filepath.Join(file, collection, "24px.svg"))
			if err != nil {
				return fault.Wrap(err)
			}
			if len(fp) != 1 {
				continue
			}
			err = col.Bundle(ctx, base, fp[0])
			if err != nil {
				return fault.Wrap(err)
			}
		}
	}
	for _, col := range collections {
		select {
		case <-done:
			return fault.Wrap(ctx.Err())
		default:
		}
		err := col.WriteFile(ctx)
		if err != nil {
			return fault.Wrap(err)
		}
	}
	return nil
}
