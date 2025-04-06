package main

import (
	"context"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"sync"

	"golang.org/x/sync/errgroup"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gosthome/icons"
	"github.com/gosthome/icons/ivg"
	_ "github.com/gosthome/icons/ivg/fortawesome/faBrands"
	_ "github.com/gosthome/icons/ivg/fortawesome/faRegular"
	_ "github.com/gosthome/icons/ivg/fortawesome/faSolid"
	_ "github.com/gosthome/icons/ivg/google/materialdesignicons"
	_ "github.com/gosthome/icons/ivg/google/materialdesigniconsoutlined"
	_ "github.com/gosthome/icons/ivg/google/materialdesigniconsround"
	_ "github.com/gosthome/icons/ivg/google/materialdesigniconssharp"
	_ "github.com/gosthome/icons/ivg/templarian/mdi"
	"github.com/majfault/signal"
	"github.com/majfault/signal/dispatcher"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.MinSize(unit.Dp(128), unit.Dp(128)))
		ui := NewUI(window)
		if err := ui.Run(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()
	app.Main()
}

type UI struct {
	sync.RWMutex
	window     *app.Window
	btn        widget.Clickable
	btnClicked signal.Signal0
	cancel     context.CancelFunc
	theme      *material.Theme
	ops        op.Ops

	allIcons icons.AllIcons
	icon     *widget.Icon
}

func NewUI(w *app.Window) *UI {
	theme := material.NewTheme()
	theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	ui := &UI{
		window:   w,
		allIcons: icons.FromCollectionKeys(ivg.Collections.Keys()),
		theme:    theme,
	}
	ui.nextIcon()
	return ui
}

func (ui *UI) frame(e app.FrameEvent) {
	ui.RLock()
	ui.RUnlock()
	// gtx is used to pass around rendering and event information.
	gtx := app.NewContext(&ui.ops, e)

	for ui.btn.Clicked(gtx) {
		ui.btnClicked.Emit()
	}

	ib := material.IconButton(ui.theme, &ui.btn, ui.icon, "Hello, Gio")
	ib.Size = unit.Dp(128)
	ib.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	ib.Layout(gtx)

	// render and handle the operations from the UI.
	e.Frame(gtx.Ops)
}

func (ui *UI) renderer() error {
	for {
		// detect the type of the event.
		switch e := ui.window.Event().(type) {
		// this is sent when the application should re-render.
		case app.FrameEvent:
			ui.frame(e)
		case app.DestroyEvent:
			ui.cancel()
			return e.Err
		}
	}
}

func (ui *UI) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	ui.cancel = cancel
	g, ctx := errgroup.WithContext(ctx)
	g.Go(ui.renderer)
	ui.btnClicked.Connect(dispatcher.Async(), ui.nextIcon)
	// g.Go(func() error {
	// 	ticker := time.NewTicker(time.Second)
	// 	done := ctx.Done()
	// 	// listen for events happening on the window.
	// 	for {
	// 		select {
	// 		case <-done:
	// 			return nil
	// 		case <-ticker.C:
	// 			ui.nextIcon()
	// 		}
	// 	}
	// })
	return g.Wait()
}

func (ui *UI) nextIcon() {
	ui.Lock()
	defer ui.Unlock()
	iconTag := ui.allIcons[rand.Intn(len(ui.allIcons))]
	fmt.Println(iconTag)
	iconResource := icons.GetResource(ivg.Collections, &iconTag)
	if iconResource == nil {
		fmt.Printf("Failed to look up %s\n", iconTag)
		return
	}
	icon, err := widget.NewIcon(iconResource.Content())
	if err != nil {
		fmt.Println(err)
		return
	}
	ui.icon = icon
	ui.window.Invalidate()
}
