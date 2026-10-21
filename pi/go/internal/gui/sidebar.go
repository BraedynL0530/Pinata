package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/BraedynL0530/Pinata/internal/core"
)

type sidebar struct {
	btn *widget.Button
	Id  int
}

func Sidebar(workspaces *WorkspaceManager, cfg *core.Config) fyne.CanvasObject {
	rect := canvas.NewRectangle(*cfg.Visual.PrimaryColor)
	rect.CornerRadius = 6

	total := workspaces.TotalWorkspaces
	for workspace := range total {
		invisibleBtn := widget.NewButton("", onClick) // wanna make these circles
		invisibleBtn.Importance = widget.LowImportance
		sidebar := &sidebar{
			btn: invisibleBtn,
			Id:  workspace,
		}
	}

	return dih
}
