package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/BraedynL0530/Pinata/internal/core"
)

type gui struct {
	cfg *core.Config
}

func setupGui(cfg *core.Config) {
	a := app.New()
	w := a.NewWindow("main")
	w.Show() // TEMP SO ERROR SHUTS UP
	home := NewHome(cfg)
	allWorkspaces := []*fyne.Container{home.Container}
	workspaces := make([]*fyne.Container, len(allWorkspaces))
	NewWorkspaceManager(workspaces)
}
