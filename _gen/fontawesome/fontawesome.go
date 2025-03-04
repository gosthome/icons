package fontawesome

import (
	"_gen/common"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func GenerateIcons() {
	mdi_dir := "../_fa.source"
	if _, err := os.Stat(mdi_dir); os.IsNotExist(err) {
		c := exec.Command("git", "clone", "-b", "6.7.2", "--depth", "1", "https://github.com/FortAwesome/Font-Awesome.git", mdi_dir)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err := c.Run()
		if err != nil {
			panic(err)
		}
	}

	os.Rename(filepath.Join(mdi_dir, "LICENSE.txt"), filepath.Join(mdi_dir, "LICENSE"))

	collections := map[string]*common.Collection{
		"brands":  common.NewColl(mdi_dir, "fortawesome/faBrands/icons.go"),
		"regular": common.NewColl(mdi_dir, "fortawesome/faRegular/icons.go"),
		"solid":   common.NewColl(mdi_dir, "fortawesome/faSolid/icons.go"),
	}
	for collection, col := range collections {
		files, err := filepath.Glob(filepath.Join(mdi_dir, "svgs", collection, "*.svg"))
		if err != nil {
			panic(err)
		}
		for _, fn := range files {
			base := filepath.Base(fn)
			base = strings.TrimSuffix(base, path.Ext(base))
			col.Bundle(base, fn)
		}
	}
	for _, col := range collections {
		col.WriteFile()
	}
}
