package ires

import (
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"

	"fmt"
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

	// Delete the expiration date image
	i.DeleteExpireImage(IMAGE_MODE_RESIZE)

	// Delete the expiration date image
	distPath := i.ImagePath(IMAGE_MODE_RESIZE)
	// When the image exists, return the image path
	//if isExistsImage(srcPath) {
	//	return i.targetImagePath(srcPath)
	//}

	inputImg, format, isExists := InputImage(i)
	if !isExists {
		return i.Uri
	}

	outputImg 		:= resize.Resize(uint(i.Width), uint(i.Height), inputImg, resize.Lanczos3)
	CreateImage(outputImg, distPath, format)
	fmt.Println(format, distPath)
	return i.TargetImagePath(distPath)
}

func (i *Ires) Crop() string {

	// Delete the expiration date image
	i.DeleteExpireImage(IMAGE_MODE_CROP)

	distPath := i.ImagePath(IMAGE_MODE_CROP)
	// When the image exists, return the image path
	//if isExistsImage(srcPath) {
	//	return i.targetImagePath(srcPath)
	//}

	inputImg, format, isImageExist := InputImage(i)
	if !isImageExist {
		return i.Uri
	}

	outputImg, _ := cutter.Crop(inputImg, cutter.Config{
		Width:  i.Width,
		Height: i.Height,
		Mode: cutter.Centered,
		Options: cutter.Copy,
	})
	CreateImage(outputImg, distPath, format)

	return i.TargetImagePath(distPath)
}

func (i *Ires) ResizeToCrop() string {

	// Delete the expiration date image
	i.DeleteExpireImage(IMAGE_MODE_RESIZE_TO_CROP)

	distPath := i.ImagePath(IMAGE_MODE_RESIZE_TO_CROP)
	// When the image exists, return the image path
	//if isExistsImage(srcPath) {
	//	return i.targetImagePath(srcPath)
	//}


	inputImg, format, isImageExist := InputImage(i)
	if !isImageExist {
		return i.Uri
	}

	outputImg := ResizeToCrop(i ,inputImg)
	CreateImage(outputImg, distPath, format)

	return i.TargetImagePath(distPath)
}
