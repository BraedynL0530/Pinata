package gui

import (
	"image/color"
	"time"

	hook "github.com/robotn/gohook"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/BraedynL0530/Pinata/internal/core"
)

type swipe struct {
	isDragging bool
	isUp       bool
	buffer     time.Duration // im a moron
}

func Sidebar(workspaces *WorkspaceManager, cfg *core.Config, swipe *swipe) *fyne.Container {
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

	if swipe.isUp && swipe.isDragging == true {
		workspaces.PrevWorkspace()
	}
	if swipe.isDragging == true && swipe.isUp != true {
		workspaces.NextWorkspace()
	}
	ui := container.NewStack(btnHolder, rect)
	return ui
}

func NewSwipe() *swipe {
	Swipe := &swipe{
		isDragging: false,
		isUp:       false,
		buffer:     1000,
	}
	return Swipe
}

func SwipeSearch(workspaces *WorkspaceManager, swipe *swipe) {
	evChan := hook.Start()
	defer hook.End()
	// useless CLANKER gosh i HATE ai especially GPT and lil G

	swipe.buffer = 400 * time.Millisecond
	for ev := range evChan {
		if ev.Kind == hook.MouseDown {
			swipe.isDragging = true //temp, channel isn't gonna work
		}
		if ev.Kind == hook.MouseUp {
			swipe.isDragging = false
		}
	}

}

//gonna stop here for the day!
