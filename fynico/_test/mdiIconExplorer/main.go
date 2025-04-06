package main

import (
	"context"
	"log/slog"
	"slices"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/hbollon/go-edlib"

	"github.com/gosthome/icons"
	"github.com/gosthome/icons/fynico"
	_ "github.com/gosthome/icons/fynico/fortawesome/faBrands"
	_ "github.com/gosthome/icons/fynico/fortawesome/faRegular"
	_ "github.com/gosthome/icons/fynico/fortawesome/faSolid"
	_ "github.com/gosthome/icons/fynico/google/materialdesignicons"
	_ "github.com/gosthome/icons/fynico/google/materialdesigniconsoutlined"
	_ "github.com/gosthome/icons/fynico/google/materialdesigniconsround"
	_ "github.com/gosthome/icons/fynico/google/materialdesigniconssharp"
	_ "github.com/gosthome/icons/fynico/templarian/mdi"
)

func main() {
	a := app.New()
	w := a.NewWindow("Icons showcase")
	iconsContainer := container.New(layout.NewAdaptiveGridLayout(3))

	iconNames := icons.FromCollectionKeys(fynico.Collections.Keys())
	keys := slices.Collect(func(yield func(string) bool) {
		for _, icon := range iconNames {
			if !yield(icon.String()) {
				return
			}
		}
	})
	e := widget.NewEntry()
	maxList := 200
	pb := widget.NewProgressBar()
	pb.Hide()
	wg := &sync.WaitGroup{}
	ctxMux := &sync.RWMutex{}
	baseCtx := context.Background()
	var ctx context.Context
	var cancel context.CancelFunc
	e.OnChanged = func(s string) {
		wg.Add(1)
		func() {
			ctxMux.RLock()
			defer ctxMux.RUnlock()
			if cancel != nil {
				cancel()
			}
		}()
		var done <-chan struct{}
		func() {
			ctxMux.Lock()
			defer ctxMux.Unlock()
			ctx, cancel = context.WithCancel(baseCtx)
			done = ctx.Done()
		}()
		go func() {
			defer wg.Done()
			iconsContainer.Hide()
			pb.Show()
			defer func() {
				iconsContainer.Show()
				pb.Hide()
			}()
			var res []string
			if s == "" {
				res = keys[:maxList]
			} else {
				var err error
				res, err = edlib.FuzzySearchSetThreshold(s, keys, maxList, 0.5, edlib.Levenshtein)
				if err != nil {
					slog.Error("Failed to search")
					return
				}
			}
			iconsContainer.RemoveAll()
			pb.Max = float64(len(res))
			for i, k := range res {
				select {
				case <-done:
					return
				default:
				}
				pb.SetValue(float64(i))
				if k == "" {
					continue
				}
				icon, _ := icons.Parse(k)
				sr := fynico.Collections.Lookup(icon.Collection, icon.Icon)
				slog.Info("found", "k", k, "for", s)
				card := widget.NewButtonWithIcon(k, theme.NewColoredResource(
					sr,
					theme.ColorNameForeground,
				), func() {
					w.Clipboard().SetContent(k)
				})
				iconsContainer.Add(card)
			}
		}()
	}
	e.SetPlaceHolder("icon name")
	scr := container.NewScroll(iconsContainer)
	scr.SetMinSize(fyne.NewSize(132, 400))
	b := container.NewBorder(e, nil, nil, nil, scr, pb)
	w.SetContent(b)
	e.OnChanged("")
	w.Resize(fyne.NewSize(1000, 500))
	w.ShowAndRun()
	wg.Wait()
}
