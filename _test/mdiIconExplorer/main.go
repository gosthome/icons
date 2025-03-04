package main

import (
	"context"
	"log/slog"
	"maps"
	"slices"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/hbollon/go-edlib"

	"github.com/gosthome/icons/templarian/mdi"
)

func main() {
	a := app.New()
	w := a.NewWindow("Icons showcase")
	iconsContainer := container.New(layout.NewAdaptiveGridLayout(3))
	keys := slices.Collect(maps.Keys(mdi.Icons))
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
				slog.Info("found", "k", k, "for", s)
				card := widget.NewButtonWithIcon(k, theme.NewColoredResource(
					mdi.Icons[k],
					theme.ColorNameForeground,
				), func() {
					w.Clipboard().SetContent("mdi:" + k)
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
