package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Home struct {
	Container *fyne.Container
	bg        *canvas.Image
}

func NewHome() *Home {
	defualtImg := canvas.NewImageFromFile(core.) // change
	defualtImg.FillMode = canvas.ImageFillContain
	c := container.NewVBox()
	return &Home{
		Container: c,
		bg:        defualtImg,
	}

}
func (home *Home) SetBG(assetPath string) *canvas.Image {
	img := canvas.NewImageFromFile(assetPath)
	img.FillMode = canvas.ImageFillContain
	return img
}
