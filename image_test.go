package image_test

import (
	"github.com/enorith/image"
	"github.com/enorith/image/assets"
	"github.com/enorith/image/elements"
	"testing"
)

func TestDraw(t *testing.T) {
	e := image.NewCanvas(512, 512, elements.Text{
		Content: "hello world 中午啊哈哈5456awdaw单位 为31 啊伟大伟大",
		Style: elements.Style{
			Top: 10,
		},
		Wrapped: true,
	}, elements.Image{
		Src: "assets/gopher.png",
		Style: elements.Style{
			Top:          100,
			BorderRadius: 24,
		},
		Width: 256,
	}).SavePNG("out.png")
	t.Error(e)
}

func TestReadFile(t *testing.T) {
	b, e := assets.ReadFile("fonts/NotoSerifCJKsc-Regular.otf")
	if e != nil {
		t.Error(e)
	}
	t.Log(len(b))
}

/**
, elements.Image{
		Src: "assets/gopher.png",
		Style: elements.Style{
			Top: 100,
		},
	}
*/
