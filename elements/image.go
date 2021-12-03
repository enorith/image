package elements

import (
	"image"
	"io"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/enorith/feather"
	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
)

type Image struct {
	Style
	Src    string
	Reader io.Reader
	Width  int
	Height int
	Limit  int
}

func (i Image) Draw(ctx *gg.Context) error {
	var im image.Image
	var e error
	if i.Reader != nil {
		im, _, e = image.Decode(i.Reader)
		if e != nil {
			logrus.Error("load image error ", e)
			return e
		}
	} else {
		if strings.HasPrefix(i.Src, "http") {
			res, e := feather.Get(i.Src)
			if e != nil {
				logrus.Error("load image error ", e)
				return e
			}
			im, _, e = image.Decode(res.Wait().Body)
			if e != nil {
				return e
			}
		} else {
			im, e = gg.LoadImage(i.Src)
			if e != nil {
				logrus.Error("load image error ", e)
				return e
			}
		}
	}

	bounds := im.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	if i.Limit != 0 && i.Width == 0 && i.Height == 0 {
		if w > h {
			i.Height = i.Limit
		} else {
			i.Width = i.Limit
		}
	}

	if i.Width != 0 || i.Height != 0 {
		im = imaging.Resize(im, i.Width, i.Height, imaging.Lanczos)
	}
	if i.Width != 0 {
		w = i.Width
		if i.Height == 0 {
			h = (i.Width / w) * h
		} else {
			h = i.Height
		}
	}
	if i.Height != 0 {
		h = i.Height
		if i.Width == 0 {
			w = (i.Height / h) * w
		} else {
			w = i.Width
		}
	}
	if i.BorderRadius != 0 {
		ctx.DrawRoundedRectangle(i.Left, i.Top, float64(w), float64(h), i.BorderRadius)
		ctx.Clip()
	}
	ctx.DrawImageAnchored(im, int(i.Left), int(i.Top), i.Anchor[0], i.Anchor[1])

	return nil
}
