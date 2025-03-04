# icons

This repository contains icons collections raedy to be used with Fyne framework.
This project is NOT affiliated with any of the source repositories.

It contains the following icons:

* Material Design Icons by Google from [google/material design icons](https://github.com/google/material-design-icons) repo.
* `@mdi/svg` npm package from [Templarian/MaterialDesign-SVG](https://github.com/Templarian/MaterialDesign-SVG) repo.
* Font-Awesome from [FortAwesome/Font-Awesome](https://github.com/FortAwesome/Font-Awesome.git) repo

All of those collections are represented as separate Go packages. There is also common convenience package `github.com/gosthome/icons`. It can be used like this:

```go

package main

import (
	// Import common package
	"github.com/gosthome/icons"

	// This imports are important as they register respective collection
	// in the common package. Collection name is the same name as package
	// `materialdesignicons:123` for icon 123 in materialdesignicons
	_ "github.com/gosthome/icons/google/materialdesignicons"
	// `materialdesigniconsoutlined:123` for icon `123` in `materialdesigniconsoutlined` collection
	_ "github.com/gosthome/icons/google/materialdesigniconsoutlined"
	// `materialdesigniconsround:123` for icon `123` in `materialdesigniconsround` collection
	_ "github.com/gosthome/icons/google/materialdesigniconsround"
	// `materialdesigniconssharp:123` for icon `123` in `materialdesigniconssharp` collection
	_ "github.com/gosthome/icons/google/materialdesigniconssharp"
	// `mdi:chip` for icon `chip` in `mdi` collection
	_ "github.com/gosthome/icons/templarian/mdi"
	// `faBrands:google` for icon `google` in `faBrands` collection
	_ "github.com/gosthome/icons/fortawesome/faBrands"
	// `faSolid:address-card` for icon address-`card` in `faRegular` collection
	_ "github.com/gosthome/icons/fortawesome/faRegular"
	// `faSolid:address-card` for icon address-`card` in `faSolid` collection
	_ "github.com/gosthome/icons/fortawesome/faSolid"
)


func main(){
	// ...
	p, err := icons.Parse(iconText)
	if err != nil {
		// handle err
		return
	}
	r := p.GetResource()
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
There is also a small test application same to MDI Icon Picker. It runs a fuzzy search on `mdi` icons and then shows appropriate icon(s).
