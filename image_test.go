package image_test

import (
	"github.com/enorith/image"
	"github.com/enorith/image/elements"
	"testing"
)

func TestDraw(t *testing.T) {
	image.NewCanvas(512, 512, elements.Text{
		Content: "hello world",
		Style: elements.Style{
			FontFile: "assets/ttf/FiraCode-Regular.ttf",
			Top:      2,
		},
		Wrapped: true,
	}).SaveJPG("out.jpg", 80)
}
