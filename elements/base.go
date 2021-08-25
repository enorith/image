package elements

import (
	"image/color"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

type Anchor [2]float64
type Box [2]int

type Style struct {
	Color           color.Color
	BackgroundColor color.Color
	FontFile        string
	FontFace        font.Face
	Top             float64
	Left            float64
	FontSize        float64
	Anchor          Anchor
	Width           float64
	Align           gg.Align
	LineSpacing     float64
	BorderRadius    float64
	Box             Box
}
