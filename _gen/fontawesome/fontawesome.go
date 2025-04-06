package fontawesome

import (
	"_gen/common"
	"context"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
)

func GenerateIcons(ctx context.Context) error {
	sourceDir := "../_fa.source"
	done := ctx.Done()
	ctx = fctx.WithMeta(ctx, "mdi_dir", sourceDir)
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		c := exec.CommandContext(ctx, "git", "clone", "-b", "6.7.2", "--depth", "1", "https://github.com/FortAwesome/Font-Awesome.git", sourceDir)
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

	os.Rename(filepath.Join(sourceDir, "LICENSE.txt"), filepath.Join(sourceDir, "LICENSE"))

	collections := map[string]*common.Collection{
		"brands":  common.NewColl(sourceDir, "fortawesome/faBrands/icons.go"),
		"regular": common.NewColl(sourceDir, "fortawesome/faRegular/icons.go"),
		"solid":   common.NewColl(sourceDir, "fortawesome/faSolid/icons.go"),
	}
	for collection, col := range collections {
		select {
		case <-done:
			return fault.Wrap(ctx.Err())
		default:
		}
		files, err := filepath.Glob(filepath.Join(sourceDir, "svgs", collection, "*.svg"))
		if err != nil {
			return fault.Wrap(err)
		}
		for _, fn := range files {
			select {
			case <-done:
				return fault.Wrap(ctx.Err())
			default:
			}
			base := filepath.Base(fn)
			base = strings.TrimSuffix(base, path.Ext(base))
			err = col.Bundle(ctx, base, fn)
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
