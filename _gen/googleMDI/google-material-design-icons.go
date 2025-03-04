package googlemdi

import (
	"_gen/common"
	"os"
	"os/exec"
	"path/filepath"
)

func GenerateIcons() {
	mdi_dir := "../_mdi.source"
	if _, err := os.Stat(mdi_dir); os.IsNotExist(err) {
		c := exec.Command("git", "clone", "--depth", "1", "https://github.com/google/material-design-icons.git", mdi_dir)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err := c.Run()
		if err != nil {
			panic(err)
		}
	}
	files, err := filepath.Glob(filepath.Join(mdi_dir, "src/*/*/"))
	if err != nil {
		panic(err)
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
			fp, err := filepath.Glob(filepath.Join(file, collection, "24px.svg"))
			if err != nil {
				panic(err)
			}
			if len(fp) != 1 {
				continue
			}
			col.Bundle(base, fp[0])
		}
	}
	for _, col := range collections {
		col.WriteFile()
	}
}
