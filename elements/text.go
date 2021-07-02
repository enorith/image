package elements

import (
	"github.com/enorith/image"
	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
)

type Text struct {
	Style
	Content string
	Wrapped bool
}

func (t Text) Draw(ctx *gg.Context) {
	if t.FontSize == 0 {
		t.FontSize = image.DefaultFontSize
	}
	if t.LineSpacing == 0 {
		t.LineSpacing = image.DefaultLineSpacing
	}
	if t.Width == 0 {
		t.Width = float64(ctx.Width())
	}
	if t.Color == nil {
		ctx.SetColor(image.DefaultColor)
	} else {
		ctx.SetColor(t.Color)
	}
	if t.FontFace != nil {
		ctx.SetFontFace(t.FontFace)
	} else if t.FontFile != "" {
		e := ctx.LoadFontFace(t.FontFile, t.FontSize)
		if e != nil {
			logrus.Error("load font error", e)
		}
	}

	if t.Wrapped {
		ctx.DrawStringWrapped(t.Content, t.Left, t.Top, t.Anchor[0], t.Anchor[1], t.Width, t.LineSpacing, t.Align)
	} else {
		ctx.DrawStringAnchored(t.Content, t.Left, t.Top, t.Anchor[0], t.Anchor[1])
	}
	ctx.Fill()
}
