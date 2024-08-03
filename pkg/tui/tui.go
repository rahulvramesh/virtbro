package tui

import (
	"github.com/rivo/tview"
)

func StartTUI() error {
	app := tview.NewApplication()

	list := tview.NewList().
		AddItem("View KVM Connections", "", 'v', nil).
		AddItem("Add Connection", "", 'a', nil).
		AddItem("Quit", "", 'q', func() {
			app.Stop()
		})

	if err := app.SetRoot(list, true).Run(); err != nil {
		return err
	}

	return nil
}
