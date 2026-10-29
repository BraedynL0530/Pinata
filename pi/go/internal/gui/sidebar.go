package gui

import (
	"image/color"

	hook "github.com/robotn/gohook"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/BraedynL0530/Pinata/internal/core"
)

type swipe struct {
	// may not use tbd...
}

func Sidebar(workspaces *WorkspaceManager, cfg *core.Config) *fyne.Container {
	rect := canvas.NewRectangle(*cfg.Visual.AccentColor)
	rect.CornerRadius = 6
	btnHolder := container.NewVBox()
	var bulletIcons []*canvas.Circle

	total := workspaces.TotalWorkspaces
	for workspace := range total {
		index := workspace

		icon := canvas.NewCircle(*cfg.Visual.TertiaryColor)
		if index == workspaces.CurrentIndex {
			icon.StrokeColor = *cfg.Visual.PrimaryColor
			icon.StrokeWidth = 4.0
		} else {
			icon.StrokeColor = color.White
			icon.StrokeWidth = 1.0
		}
		bulletIcons = append(bulletIcons, icon)

		btn := widget.NewButton("", func() {
			workspaces.SetWorkspace(index)
		})
		btn.Importance = widget.LowImportance

		trueButton := container.NewStack(icon, btn)
		btnHolder.Add(trueButton)
	}

	workspaces.OnChanged = func(activeIndex int) {
		for i, circle := range bulletIcons {
			if i == activeIndex {
				circle.StrokeColor = *cfg.Visual.PrimaryColor
				circle.StrokeWidth = 4.0
			} else {
				circle.StrokeWidth = 1.0
			}
			circle.Refresh()
			//animation here
		}
	}

	ui := container.NewStack(btnHolder, rect)
	return ui
}

func SwipeSearch(workspaces *WorkspaceManager, cfg *core.Config) { // autocomplete but likely wont need cfg
	evChan := hook.Start()
	defer hook.End()

}
