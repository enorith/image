package image_test

import (
	"testing"

	"github.com/enorith/image"
	"github.com/enorith/image/assets"
	"github.com/enorith/image/elements"
)

func TestDraw(t *testing.T) {
	e := image.NewCanvas(512, 512, elements.Text{
		Content: "hello world 你好，中文文字，滚滚长江东逝水，浪花淘尽英雄。是非成败转头空：青山依旧在，几度夕阳红。白发渔樵江渚上，惯看秋月春风。一壶浊酒喜相逢：古今多少事，都付笑谈中。话说天下大势，分久必合，合久必分：周末七国分争，并入于秦。",
		Style: elements.Style{
			Top:         320,
			Left:        12,
			Width:       488,
			LineSpacing: 1,
		},
		Wrapped: true,
	}, elements.Image{
		Src: "https://picsum.photos/600/300",
		Style: elements.Style{
			Top:          12,
			BorderRadius: 12,
			Left:         12,
		},
	}).SavePNG("out.png")
	if e != nil {
		t.Error(e)
	}
}

func TestLimit(t *testing.T) {
	w := 100

	e := image.NewCanvas(200, 200, elements.Image{
		Src:   "https://picsum.photos/600/300",
		Limit: w,
	}).SavePNG("out_limit.png")
	if e != nil {
		t.Error(e)
	}
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
