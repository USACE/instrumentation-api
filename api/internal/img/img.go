package img

import (
	"image"
	"image/png"
	"io"

	"golang.org/x/image/draw"
)

func Resize(r io.Reader, w io.Writer, limit image.Rectangle) error {
	src, err := png.Decode(r)
	if err != nil {
		return err
	}

	maxX := src.Bounds().Max.X
	maxY := src.Bounds().Max.Y

	xLimit := limit.Max.X
	yLimit := limit.Max.Y

	if maxX <= xLimit && maxY <= yLimit {
		return png.Encode(w, src)
	}

	scalingFactor := 1
	if maxX > xLimit && (maxX > maxY || maxX == maxY) {
		scalingFactor = xLimit / maxX
	} else if maxY > yLimit && (maxY > maxX || maxX == maxY) {
		scalingFactor = yLimit / maxY
	}

	xLimit *= scalingFactor
	yLimit *= scalingFactor

	dst := image.NewRGBA(image.Rect(0, 0, xLimit, yLimit))
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	return png.Encode(w, dst)
}
