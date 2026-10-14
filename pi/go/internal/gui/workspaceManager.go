package gui

import "fyne.io/fyne/v2"

//like the scroll bar that switches from the home menu w/ acsii terminal & weather -> apps -> webcam -> whatever else

type WorkspaceManager struct {
	CurrentIndex    int // which workspace is selected
	TotalWorkspaces int
	Workspaces      []*fyne.Container
	OnChanged       func(index int)
}

func NewWorkspaceManager(workspaces []*fyne.Container) *WorkspaceManager {
	return &WorkspaceManager{
		CurrentIndex:    0,
		TotalWorkspaces: len(workspaces),
		Workspaces:      workspaces,
	}
}

func (wm *WorkspaceManager) NextWorkspace() {
	if wm.CurrentIndex < wm.TotalWorkspaces-1 {
		wm.CurrentIndex++
	} else {
		wm.CurrentIndex = 0
	}
	wm.UpdateVisibility()
}

func (wm *WorkspaceManager) PrevWorkspace() {
	if wm.CurrentIndex > 0 {
		wm.CurrentIndex--
	} else {
		wm.CurrentIndex = wm.TotalWorkspaces - 1
	}
	wm.UpdateVisibility()
}

func (wm *WorkspaceManager) SetWorkspace(index int) {
	if index >= 0 && index < wm.TotalWorkspaces {
		wm.CurrentIndex = index
		wm.UpdateVisibility()
	}
}

func (wm *WorkspaceManager) UpdateVisibility() {
	for i, ws := range wm.Workspaces {
		if i == wm.CurrentIndex {
			ws.Show()
		} else {
			ws.Hide()
		}
	}
}
