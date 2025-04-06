package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/gosthome/icons"
	"github.com/gosthome/icons/fyneicons"
	_ "github.com/gosthome/icons/fyneicons/fortawesome/faBrands"
	_ "github.com/gosthome/icons/fyneicons/fortawesome/faRegular"
	_ "github.com/gosthome/icons/fyneicons/fortawesome/faSolid"
	_ "github.com/gosthome/icons/fyneicons/google/materialdesignicons"
	_ "github.com/gosthome/icons/fyneicons/google/materialdesigniconsoutlined"
	_ "github.com/gosthome/icons/fyneicons/google/materialdesigniconsround"
	_ "github.com/gosthome/icons/fyneicons/google/materialdesigniconssharp"
	_ "github.com/gosthome/icons/fyneicons/templarian/mdi"
)

func main() {
	resNotFound := widget.NewLabel("resource not found")
	resNotFound.Wrapping = fyne.TextWrapWord
	resNotFound.Importance = widget.WarningImportance
	l := widget.NewLabel("")
	l.Wrapping = fyne.TextWrapWord
	l.Importance = widget.WarningImportance

	a := app.New()
	w := a.NewWindow("Icons parser")
	iconsContainer := container.NewGridWrap(fyne.NewSize(256, 256))
	e := widget.NewEntry()
	e.OnChanged = func(s string) {
		iconsContainer.RemoveAll()
		p, err := icons.Parse(s)
		if err != nil {
			l.SetText(err.Error())
			iconsContainer.Add(l)
			return
		}
		r := icons.GetResource(fyneicons.Collections, p)
		if r == nil {
			iconsContainer.Add(resNotFound)
			return
		}
		iconsContainer.Add(canvas.NewImageFromResource(
			theme.NewColoredResource(
				r,
				theme.ColorNameForeground,
			)))
	}

	b := container.NewBorder(e, nil, nil, nil, iconsContainer)
	w.SetContent(b)
	w.Resize(fyne.NewSize(300, 256))
	w.ShowAndRun()
}
