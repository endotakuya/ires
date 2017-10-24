package ires

import (
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

const (
	IMAGE_MODE_RESIZE int = iota
	IMAGE_MODE_CROP
	IMAGE_MODE_RESIZE_TO_CROP
	IMAGE_MODE_ORIGINAL
)

type Size struct {
	Width, Height int
}

type Ires struct {
	Uri 		string
	Size
	Dir 		string
	Expire 		string
	IsLocal 	bool
}


func (i *Ires) Resize() string {
	// Check image type
	i.isLocalFile()

	// Delete the expiration date image
	i.deleteExpireImage(IMAGE_MODE_RESIZE)

	distPath := i.imagePath(IMAGE_MODE_RESIZE)
	// When the image exists, return the image path
	if isExistsImage(distPath) {
		return i.targetImagePath(distPath)
	}

	inputImg, format, isImageExist := inputImage(i)
	if !isImageExist {
		return i.Uri
	}

	outputImg := resize.Resize(uint(i.Width), uint(i.Height), inputImg, resize.Lanczos3)
	createImage(outputImg, distPath, format)
	return i.targetImagePath(distPath)
}

func (i *Ires) Crop() string {
	// Check image type
	i.isLocalFile()

	// Delete the expiration date image
	i.deleteExpireImage(IMAGE_MODE_CROP)

	distPath := i.imagePath(IMAGE_MODE_CROP)
	// When the image exists, return the image path
	if isExistsImage(distPath) {
		return i.targetImagePath(distPath)
	}

	inputImg, format, isImageExist := inputImage(i)
	if !isImageExist {
		return i.Uri
	}

	outputImg, _ := cutter.Crop(inputImg, cutter.Config{
		Width:  i.Width,
		Height: i.Height,
		Mode: cutter.Centered,
		Options: cutter.Copy,
	})
	createImage(outputImg, distPath, format)

	return i.targetImagePath(distPath)
}

func (i *Ires) ResizeToCrop() string {
	// Check image type
	i.isLocalFile()

	// Delete the expiration date image
	i.deleteExpireImage(IMAGE_MODE_RESIZE_TO_CROP)

	distPath := i.imagePath(IMAGE_MODE_RESIZE_TO_CROP)
	// When the image exists, return the image path
	if isExistsImage(distPath) {
		return i.targetImagePath(distPath)
	}

	inputImg, format, isImageExist := inputImage(i)
	if !isImageExist {
		return i.Uri
	}

	outputImg := resizeToCrop(i ,inputImg)
	createImage(outputImg, distPath, format)

	return i.targetImagePath(distPath)
}
