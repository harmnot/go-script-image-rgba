package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"strconv"
	"path/filepath"
)

func AdjustImageColors(img image.Image, percentage int) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	// Map the percentage to 0-255
	value := percentage * 255 / 100

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			// Ensure we correctly handle each pixel's RGB channels
			oldColor := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			if value >= 0 {
				// Increase red channel
				oldColor.R = uint8(min(int(oldColor.R)+value, 255))
			} else {
				// Increase blue channel
				oldColor.B = uint8(min(int(oldColor.B)-value, 255))
			}
			newImg.Set(x, y, oldColor)
		}
	}
	return newImg
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Read image from file
	imgPath := os.Args[1]
	percentage, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("You need to provide an integer value as the second argument, originalError: " + err.Error() + ".")
	}

	file, err := os.Open(imgPath)
	if err != nil {
	    panic(`File not found, please provide a valid file path, originalError: ` + err.Error() + `.`)
	}
	defer file.Close()

	_, formatInfo, err := image.DecodeConfig(file)
	if err != nil {
	    panic(`Could not decode image, originalError: ` + err.Error() + `.`)
    }

    if filepath.Ext(imgPath) != ".jpeg" || formatInfo != "jpeg" {
       panic("File is not a jpeg, please provide a jpeg file, originalError: " + err.Error() + ".")
    }

    // By using file.Seek(0, 0), you reset the file pointer to the beginning of the file,
    // ensuring that jpeg.Decode(file) can read and decode the entire image correctly.
    // offset the file pointer to the beginning of the file
    // whence 0 means relative to the origin of the file
    file.Seek(0, 0)
    // if not use the offset, the image will not be decoded
    // because the file pointer is at the end of the file
    // also will get the error " missing SOI marker"

	// Decode the image
	img, err := jpeg.Decode(file)
	if err != nil {
		panic(`Could not decode image, originalError: ` + err.Error() + `.`)
	}

	// Adjust image colors
	newImg := AdjustImageColors(img, percentage)

	// Write new image to file
	outFile, err := os.Create("./images/output.jpeg")
	if err != nil {
		panic(`Could not create file, please check your permissions, originalError: ` + err.Error() + `.`)
	}
	defer outFile.Close()

	// Encode and save new image
	err = jpeg.Encode(outFile, newImg, &jpeg.Options{Quality: 100})
	if err != nil {
		panic(`Could not encode image, originalError: ` + err.Error() + `.`)
	}
}