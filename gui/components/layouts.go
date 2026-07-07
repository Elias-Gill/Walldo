package components

import "fyne.io/fyne/v2"

const padding float32 = 4

type PriorityLayout struct {
	PrimaryRatio float32
}

func (l *PriorityLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) == 0 {
		return
	}
	if l.PrimaryRatio <= 0 {
		l.PrimaryRatio = 0.85
	}

	primarySize := size.Width * l.PrimaryRatio
	remaining := len(objects) - 1
	secondarySize := (size.Width - primarySize) / float32(remaining)

	objects[0].Resize(fyne.Size{Width: primarySize, Height: size.Height})
	objects[0].Move(fyne.Position{X: 0, Y: 0})

	xPos := primarySize
	for i := 1; i < len(objects); i++ {
		objects[i].Resize(fyne.Size{Width: secondarySize, Height: size.Height})
		objects[i].Move(fyne.Position{X: xPos + padding, Y: 0})
		xPos += secondarySize
	}
}

func (l *PriorityLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) == 0 {
		return fyne.Size{Width: 0, Height: 0}
	}

	primaryMin := objects[0].MinSize()
	var maxSecHeight float32
	var totalSecWidth float32

	for i := 1; i < len(objects); i++ {
		min := objects[i].MinSize()
		totalSecWidth += min.Width
		if min.Height > maxSecHeight {
			maxSecHeight = min.Height
		}
	}

	return fyne.Size{
		Width:  primaryMin.Width + totalSecWidth + float32(len(objects)-1)*padding,
		Height: max(primaryMin.Height, maxSecHeight),
	}
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
