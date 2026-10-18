package gui

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func setupGui() {
	//primaryColor := config here!
	a := app.New()
	w := a.NewWindow("main")
	w.SetMaster()                       // background, and workspaces, apps will be their own window? or maybe panells havent decied
	homeWorkspace := container.NewVBox( // also may need to change the type of container in .NewH/Vbox!
	//temp (bg, app luancher
	)
	wacomWorkspace := container.NewHBox(
	//placeholder
	)
	webcamWorkspace := container.NewVBox(
	//yet again
	)

	settingsWorkspace := container.NewHBox(
	//temp
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("Home", homeWorkspace), // probally gonna just make them () circles vertically stacked + part of main screen on right and center if possible
		container.NewTabItem("Settings", settingsWorkspace),
		container.NewTabItem("Webcam", webcamWorkspace),
		container.NewTabItem("Drawing", wacomWorkspace),
	)
	fmt.Println(tabs)
}
