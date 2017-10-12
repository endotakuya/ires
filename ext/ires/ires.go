package ires

import (
	"strings"

	"github.com/endotakuya/ires/ext/operate"
	"github.com/endotakuya/ires/ext/util/uri"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

const (
	IMAGE_MODE_RESIZE int = iota
	IMAGE_MODE_CROP
	IMAGE_MODE_RESIZE_TO_CROP
	IMAGE_MODE_ORIGINAL
)

type Request struct {
	Uri 	string
	Width 	int
	Height 	int
	Dir 	string
	Expire 	string
}


func (r *Request) Resize() string {
	size := []int{r.Width, r.Height}

	// Delete the expiration date image
	util.DeleteExpireImage(r.Uri, r.Dir, IMAGE_MODE_RESIZE, size...)

	path := r.fullPath(IMAGE_MODE_RESIZE)
	// When the image exists, return the image path
	if util.IsExistsImage(path) {
		return r.imagePath(path)
	}

	inputImg, _, isImageExist := operate.InputImage(r.Uri, r.originalFullPath())
	if !isImageExist {
		return r.Uri
	}

	outputImg 		:= resize.Resize(uint(r.Width), uint(r.Height), inputImg, resize.Lanczos3)
	_, fullPath, _ 	:= operate.CreateImage(outputImg, path)

	return r.imagePath(fullPath)
}

func (r *Request) Crop() string {
	size := []int{r.Width, r.Height}

	// Delete the expiration date image
	util.DeleteExpireImage(r.Uri, r.Dir, IMAGE_MODE_CROP, size...)

	path := r.fullPath(IMAGE_MODE_CROP)
	// When the image exists, return the image path
	if util.IsExistsImage(path) {
		return r.imagePath(path)
	}

	inputImg, _, isImageExist := operate.InputImage(r.Uri, r.originalFullPath())
	if !isImageExist {
		return r.Uri
	}

	outputImg, _ := cutter.Crop(inputImg, cutter.Config{
		Width:  r.Width,
		Height: r.Height,
		Mode: cutter.Centered,
		Options: cutter.Copy,
	})
	_, fullPath, _ := operate.CreateImage(outputImg, path)

	return r.imagePath(fullPath)
}

func (r *Request) ResizeToCrop() string {
	size := []int{r.Width, r.Height}

	// Delete the expiration date image
	util.DeleteExpireImage(r.Uri, r.Dir, IMAGE_MODE_RESIZE_TO_CROP, size...)

	path := r.fullPath(IMAGE_MODE_RESIZE_TO_CROP)
	// When the image exists, return the image path
	if util.IsExistsImage(path) {
		return r.imagePath(path)
	}

	inputImg, imgPath, isImageExist := operate.InputImage(r.Uri, r.originalFullPath())
	if !isImageExist {
		return r.Uri
	}

	outputImg 		:= operate.ResizeToCrop(imgPath, size, inputImg)
	_, fullPath, _ 	:= operate.CreateImage(outputImg, path)

	return r.imagePath(fullPath)
}

func (r *Request) fullPath(mode int) string {
	size := []int{ r.Width, r.Height }
	return util.NewImagePath(r.Uri, r.Dir, r.Expire, mode, size...)
}

func (r *Request) originalFullPath() string {
	return util.NewImagePath(r.Uri, r.Dir, r.Expire, IMAGE_MODE_ORIGINAL)
}

func (r *Request) imagePath(fullPath string) string {
	return strings.Replace(fullPath, r.Dir, "", -1)
}