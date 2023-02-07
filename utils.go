package main

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"net/http"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

// get minimum of two numbers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// get type(format) of input data
func getDataType(imageBytes *[]byte) string {
	return http.DetectContentType(*imageBytes)
}

// convert jpeg image into png
func jpegToPng(imageBytes []byte) []byte {
	pngEncoder := &png.Encoder{
		CompressionLevel: png.BestCompression,
	}

	// get image format
	dataType := getDataType(&imageBytes)

	if dataType == "image/jpeg" {
		// decode jpeg image into image type
		srcImage, err := jpeg.Decode(bytes.NewReader(imageBytes))
		handleError(err)

		buffer := new(bytes.Buffer)

		// encode image type into png
		err = pngEncoder.Encode(buffer, srcImage)
		handleError(err)

		return buffer.Bytes()
	} else {
		panic("file is not image/jpeg")
	}
}

// convert png image into jpeg
func pngToJpeg(imageBytes []byte) []byte {
	// get image format
	dataType := getDataType(&imageBytes)

	if dataType == "image/png" {
		// decode png image into image type
		srcImage, err := png.Decode(bytes.NewReader(imageBytes))
		handleError(err)

		buffer := new(bytes.Buffer)

		// encode image type into jpeg
		err = jpeg.Encode(buffer, srcImage, nil)
		handleError(err)

		return buffer.Bytes()
	} else {
		panic("file is not image/png")
	}
}

// cut jpeg or png images to square (same width and height)
func toSquare(imageBytes []byte) []byte {
	// get image format
	dataType := getDataType(&imageBytes)

	var srcImage image.Image
	var destImage draw.Image
	var rect image.Rectangle
	var err error
	buffer := new(bytes.Buffer)

	if dataType == "image/png" {
		// decode png image into image type
		srcImage, err = png.Decode(bytes.NewReader(imageBytes))
		handleError(err)
	} else if dataType == "image/jpeg" {
		// decode jpeg image into image type
		srcImage, err = jpeg.Decode(bytes.NewReader(imageBytes))
		handleError(err)
	} else {
		panic("image type not supported")
	}

	// get min of width and height of image
	min := min(srcImage.Bounds().Max.X, srcImage.Bounds().Max.Y)

	// create empty rectangle for square image
	rect = image.Rectangle{image.Pt(0, 0), image.Pt(min, min)}

	// create empty image with the dimentions of rectangle
	destImage = image.NewRGBA(rect)

	// draw src image on the empty square image
	draw.Draw(destImage, rect, srcImage, rect.Min, draw.Src)

	if dataType == "image/png" {
		// encode image type back into png
		err = png.Encode(buffer, destImage)
		handleError(err)

		return buffer.Bytes()
	} else {
		// encode image type back into jpeg
		err = jpeg.Encode(buffer, destImage, nil)
		handleError(err)

		return buffer.Bytes()
	}
}
