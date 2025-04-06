# icons

This repository contains icons collections ready to be used with [Fyne framework](https://fyne.io) and as icons in IconVG format, suitable for [GIO framework](https://gioui.org/).
This project is NOT affiliated with any of the source repositories.

It contains the following icons:

* Material Design Icons by Google from [google/material design icons](https://github.com/google/material-design-icons) repo.
* `@mdi/svg` npm package from [Templarian/MaterialDesign-SVG](https://github.com/Templarian/MaterialDesign-SVG) repo.
* Font-Awesome from [FortAwesome/Font-Awesome](https://github.com/FortAwesome/Font-Awesome.git) repo

All of those collections are represented as separate Go packages:
* `github.com/gosthome/icons` with convenience functions to parse icon names like `mdi:chip`.
* `github.com/gosthome/icons/fynico` and `github.com/gosthome/icons/ivg` store icon collections for `fyne` (in svg) and IconVG respectively.
* `github.com/gosthome/icons/fynico/google/materialdesignicons`, ... (see below for more ) store icon collections with their respective licenses.

```go

package main

import (
	// Import common package
	"github.com/gosthome/icons"
	"github.com/gosthome/icons/fynico"

	// This imports are important as they register respective collection
	// in the common package. Collection name is the same name as package
	// `materialdesignicons:123` for icon 123 in materialdesignicons
	_ "github.com/gosthome/icons/fynico/google/materialdesignicons"
	// `materialdesigniconsoutlined:123` for icon `123` in `materialdesigniconsoutlined` collection
	_ "github.com/gosthome/icons/fynico/google/materialdesigniconsoutlined"
	// `materialdesigniconsround:123` for icon `123` in `materialdesigniconsround` collection
	_ "github.com/gosthome/icons/fynico/google/materialdesigniconsround"
	// `materialdesigniconssharp:123` for icon `123` in `materialdesigniconssharp` collection
	_ "github.com/gosthome/icons/fynico/google/materialdesigniconssharp"
	// `mdi:chip` for icon `chip` in `mdi` collection
	_ "github.com/gosthome/icons/fynico/fyneicons/templarian/mdi"
	// `faBrands:google` for icon `google` in `faBrands` collection
	_ "github.com/gosthome/icons/fynico/fortawesome/faBrands"
	// `faSolid:address-card` for icon address-`card` in `faRegular` collection
	_ "github.com/gosthome/icons/fynico/fortawesome/faRegular"
	// `faSolid:address-card` for icon address-`card` in `faSolid` collection
	_ "github.com/gosthome/icons/fynico/fortawesome/faSolid"
)


func main(){
	// ...
	p, err := icons.Parse(iconText)
	if err != nil {
		// handle err
		return
	}
	r := fynico.Collections.Lookup(p.Collection, p.Icon)
	if r == nil {
		// handle err
		return
	}
	icon := theme.NewColoredResource(
			r,
			theme.ColorNameForeground,
	)
	// use icon
	// ...
}

```

See the full working example of icon parsing [here](_test/iconParser/main.go).
There is also a small test application same to MDI Icon Picker. It runs a fuzzy search on all included icons and then shows appropriate icon(s).
