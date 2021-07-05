package elements

import (
	"github.com/enorith/image"
	"github.com/enorith/image/assets"
	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font/opentype"
	"io/ioutil"
	"strings"
)

type Text struct {
	Style
	Content string
	Wrapped bool
}

func (t Text) Draw(ctx *gg.Context) error {
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
		e := loadFont(ctx, t.FontFile, t.FontSize)
		if e != nil {
			logrus.Error("load font error ", e)
			return e

		}
	} else {
		e := loadDefaultFont(ctx, t.FontSize)
		if e != nil {
			logrus.Error("load font error ", e)
			return e
		}
	}
	if t.Wrapped {
		ctx.DrawStringWrapped(t.Content, t.Left, t.Top, t.Anchor[0], t.Anchor[1], t.Width, t.LineSpacing, t.Align)
	} else {
		ctx.DrawStringAnchored(t.Content, t.Left, t.Top, t.Anchor[0], t.Anchor[1])
	}
	ctx.Fill()
	return nil
}

func loadDefaultFont(ctx *gg.Context, points float64) error {
	b, e := assets.ReadFile("fonts/" + image.DefaultFont)
	if e != nil {
		return e
	}

	return setFontOpenType(ctx, b, points)
}

func loadFont(ctx *gg.Context, fontPath string, points float64) error {
	if strings.HasSuffix(fontPath, "otf") {
		b, e := ioutil.ReadFile(fontPath)
		if e != nil {
			return e
		}

		err := setFontOpenType(ctx, b, points)
		if err != nil {
			return err
		}
		return nil
	} else {
		return ctx.LoadFontFace(fontPath, points)
	}
}

func setFontOpenType(ctx *gg.Context, b []byte, points float64) error {
	f, e := opentype.Parse(b)
	if e != nil {
		return e
	}
	face, e := opentype.NewFace(f, &opentype.FaceOptions{Size: points, DPI: 72})
	if e != nil {
		return e
	}
	ctx.SetFontFace(face)
	return nil
}
