package templarianmaterialdesignsvg

import (
	"_gen/common"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func GenerateIcons() {
	mdi_dir := "../_mdi_svg.source"
	if _, err := os.Stat(mdi_dir); os.IsNotExist(err) {
		c := exec.Command("git", "clone", "-b", "v7.4.47", "--depth", "1", "https://github.com/Templarian/MaterialDesign-SVG.git", mdi_dir)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err := c.Run()
		if err != nil {
			panic(err)
		}
	}
	files, err := filepath.Glob(filepath.Join(mdi_dir, "svg/*.svg"))
	if err != nil {
		panic(err)
	}
	col := common.NewColl(mdi_dir, "templarian/mdi/icons.go")
	for _, fn := range files {
		base := filepath.Base(fn)
		base = strings.TrimSuffix(base, path.Ext(base))
		col.Bundle(base, fn)
	}
	col.WriteFile()
}
