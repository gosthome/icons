package templarianmaterialdesignsvg

import (
	"_gen/common"
	"context"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/Southclaws/fault"
)

func GenerateIcons(ctx context.Context) error {
	done := ctx.Done()
	mdi_dir := "../_mdi_svg.source"
	if _, err := os.Stat(mdi_dir); os.IsNotExist(err) {
		c := exec.Command("git", "clone", "-b", "v7.4.47", "--depth", "1", "https://github.com/Templarian/MaterialDesign-SVG.git", mdi_dir)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
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
	files, err := filepath.Glob(filepath.Join(mdi_dir, "svg/*.svg"))
	if err != nil {
		return fault.Wrap(err)
	}
	col := common.NewColl(mdi_dir, "templarian/mdi/icons.go")
	for _, fn := range files {
		select {
		case <-done:
			return fault.Wrap(ctx.Err())
		default:
		}
		base := filepath.Base(fn)
		base = strings.TrimSuffix(base, path.Ext(base))
		col.Bundle(ctx, base, fn)
	}
	err = col.WriteFile(ctx)
	if err != nil {
		return fault.Wrap(err)
	}
	return nil
}
