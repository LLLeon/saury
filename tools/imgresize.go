package tools

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"strings"

	"github.com/nfnt/resize"
)

const (
	formatJPG  = "jpg"
	formatJPEG = "jpeg"
	formatPNG  = "png"
)

// ImageResize passes in the image base64 format string and compression ratio,
// and returns the compressed image base64 format string.
func ImageResize(srcBinData string, ratio float64) ([]byte, error) {
	var (
		imgByte []byte
	)

	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(srcBinData))

	img, format, err := image.Decode(dec)
	if err != nil {
		return nil, err
	}

	b := img.Bounds()
	srcWidth := b.Max.X
	srcHeight := b.Max.Y
	dstWidth, dstHeight := CalculateDstSize(srcWidth, srcHeight, ratio)

	// compressed image
	dstImg := resize.Resize(uint(dstWidth), uint(dstHeight), img, resize.Lanczos3)

	// convert Image to []byte
	if format == formatJPEG || format == formatJPG {
		buf := new(bytes.Buffer)
		err := jpeg.Encode(buf, dstImg, nil)
		if err != nil {
			return nil, err
		}

		imgByte = buf.Bytes()
	} else if format == formatPNG {
		buf := new(bytes.Buffer)
		err := png.Encode(buf, dstImg)
		if err != nil {
			return nil, err
		}

		imgByte = buf.Bytes()
	}

	return imgByte, nil
}

func CalculateDstSize(srcWidth, srcHeight int, compressRatio float64) (int, int) {
	dstWidth := int(math.Ceil(float64(srcWidth) * compressRatio))
	dstHeight := int(math.Ceil(float64(srcHeight) * compressRatio))
	return dstWidth, dstHeight
}
