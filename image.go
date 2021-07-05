package image

import (
	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
	"image/color"
)

var DefaultColor = color.Black
var DefaultBackgroundColor = color.White
var DefaultFontSize float64 = 16
var DefaultLineSpacing float64 = 1
var DefaultFont = "NotoSerifCJKsc-Regular.otf"

type Element interface {
	Draw(ctx *gg.Context) error
}

type Canvas struct {
	*gg.Context
	elements        []Element
	draw            bool
	backgroundColor color.Color
}

func (c *Canvas) SetBackgroundColor(color color.Color) *Canvas {
	c.backgroundColor = color
	return c
}

func (c *Canvas) Draw() *Canvas {
	if c.draw == true {
		return c
	}
	c.SetColor(c.backgroundColor)
	c.Clear()
	for _, e := range c.elements {
		err := e.Draw(c.Context)
		if err != nil {
			logrus.Error(err)
		}
	}
	c.draw = true
	return c
}

// SavePNG encodes the image as a PNG and writes it to disk.
func (c *Canvas) SavePNG(path string) error {
	c.Draw()
	return c.Context.SavePNG(path)
}

// SaveJPG encodes the image as a JPG and writes it to disk.
func (c *Canvas) SaveJPG(path string, quality int) error {
	c.Draw()
	return c.Context.SaveJPG(path, quality)
}

func NewCanvas(w int, h int, elements ...Element) *Canvas {
	ctx := gg.NewContext(w, h)

	return &Canvas{
		Context:         ctx,
		elements:        elements,
		backgroundColor: DefaultBackgroundColor,
	}
}
