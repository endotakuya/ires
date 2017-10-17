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

	// Delete the expiration date image
	//util.DeleteExpireImage(i.Uri, i.Dir, IMAGE_MODE_CROP, size...)

	// Delete the expiration date image
	srcPath := i.imagePath(IMAGE_MODE_RESIZE)
	// When the image exists, return the image path
	//if isExistsImage(srcPath) {
	//	return i.targetImagePath(srcPath)
	//}

	inputImg, format, isExists := inputImage(i)
	if !isExists {
		return i.Uri
	}

	outputImg 		:= resize.Resize(uint(i.Width), uint(i.Height), inputImg, resize.Lanczos3)
	_, distPath, _ 	:= createImage(outputImg, srcPath, format)

	return i.targetImagePath(distPath)
}

func (i *Ires) Crop() string {

	// Delete the expiration date image
	//util.DeleteExpireImage(i.Uri, i.Dir, IMAGE_MODE_CROP, size...)

	srcPath := i.imagePath(IMAGE_MODE_CROP)
	// When the image exists, return the image path
	//if isExistsImage(srcPath) {
	//	return i.targetImagePath(srcPath)
	//}

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
	_, distPath, _ := createImage(outputImg, srcPath, format)

	return i.targetImagePath(distPath)
}

func (i *Ires) ResizeToCrop() string {

	// Delete the expiration date image
	//util.DeleteExpireImage(i.Uri, i.Dir, IMAGE_MODE_RESIZE_TO_CROP, size...)

	srcPath := i.imagePath(IMAGE_MODE_RESIZE_TO_CROP)
	// When the image exists, return the image path
	//if isExistsImage(srcPath) {
	//	return i.targetImagePath(srcPath)
	//}


	inputImg, format, isImageExist := inputImage(i)
	if !isImageExist {
		return i.Uri
	}

	outputImg 		:= resizeToCrop(i ,inputImg)
	_, fullPath, _ 	:= createImage(outputImg, srcPath, format)

	return i.targetImagePath(fullPath)
}
