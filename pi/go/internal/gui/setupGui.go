package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func setupGui() {
	//primaryColor := config here!
	a := app.New()
	w := a.NewWindow("main")
	workspaces :=make([]*fyne.Container, Home.Container)
	NewWorkspaceManager(workspaces)