package img

import (
	"image"
	"image/jpeg"
	_ "image/png"
	"io"

	"golang.org/x/image/draw"
)

var jpegOpts = &jpeg.Options{Quality: 75}

func Resize(r io.Reader, w io.Writer, limit image.Rectangle) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	boundsX := src.Bounds().Max.X
	boundsY := src.Bounds().Max.Y

	xLimit := limit.Max.X
	yLimit := limit.Max.Y

	if boundsX <= xLimit && boundsY <= yLimit {
		return jpeg.Encode(w, src, jpegOpts)
	}

	// maintain aspect ratio during resize
	scalingFactor := 1
	if boundsX > xLimit && (boundsX > boundsY || boundsX == boundsY) {
		scalingFactor = xLimit / boundsX
	} else if boundsY > yLimit && (boundsY > boundsX || boundsX == boundsY) {
		scalingFactor = yLimit / boundsY
	}

	boundsX *= scalingFactor
	boundsY *= scalingFactor

	dst := image.NewRGBA(image.Rect(0, 0, boundsX, boundsY))
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	return jpeg.Encode(w, dst, jpegOpts)
}
