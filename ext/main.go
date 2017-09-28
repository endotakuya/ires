package main

import (
	"C"
	"strings"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
	"github.com/endotakuya/ires/ext/util/uri"
	"github.com/endotakuya/ires/ext/operate"
)

const (
	IMAGE_MODE_RESIZE int = iota
	IMAGE_MODE_CROP
	IMAGE_MODE_RESIZE_TO_CROP
	IMAGE_MODE_ORIGINAL
)

func init() {}
func main() {}

//export resizeImage
func resizeImage(Uri *C.char, width, height int, Dir, Expire *C.char) *C.char {
	uri		:= C.GoString(Uri)
	dir 	:= C.GoString(Dir)
	expire	:= C.GoString(Expire)

	size := []int{width, height}
	path := util.NewImagePath(uri, dir, expire, IMAGE_MODE_RESIZE, size...)
	originalPath := util.NewImagePath(uri, dir, expire, IMAGE_MODE_ORIGINAL)

	// When the image exists, return the image path
	if !util.IsEmptyImage(path) {
		return C.CString(strings.Replace(path, dir, "", -1))
	}

	inputImg, _, isImageExist := operate.InputImage(uri, originalPath)
	if !isImageExist {
		return C.CString(uri)
	}
	outputImg 	:= resize.Resize(uint(width), uint(height), inputImg, resize.Lanczos3)

	_, filePath, _ := operate.CreateImage(outputImg, path)

	fileName := strings.Replace(filePath, dir, "", -1)
	return C.CString(fileName)
}

//export cropImage
func cropImage(Uri *C.char, width, height int, Dir, Expire *C.char) *C.char {
	uri		:= C.GoString(Uri)
	dir 	:= C.GoString(Dir)
	expire	:= C.GoString(Expire)

	size := []int{width, height}
	path := util.NewImagePath(uri, dir, expire, IMAGE_MODE_CROP, size...)
	originalPath := util.NewImagePath(uri, dir, expire, IMAGE_MODE_ORIGINAL)

	// When the image exists, return the image path
	if !util.IsEmptyImage(path) {
		return C.CString(strings.Replace(path, dir, "", -1))
	}

	inputImg, _, isImageExist := operate.InputImage(uri, originalPath)
	if !isImageExist {
		return C.CString(uri)
	}
	outputImg, _ := cutter.Crop(inputImg, cutter.Config{
		Width:  width,
		Height: height,
		Mode: cutter.Centered,
		Options: cutter.Copy,
	})

	_, filePath, _ := operate.CreateImage(outputImg, path)

	fileName := strings.Replace(filePath, dir, "", -1)
	return C.CString(fileName)
}

//export resizeToCropImage
func resizeToCropImage(Uri *C.char, width, height int, Dir, Expire *C.char) *C.char {
	uri		:= C.GoString(Uri)
	dir 	:= C.GoString(Dir)
	expire	:= C.GoString(Expire)

	size := []int{width, height}
	path := util.NewImagePath(uri, dir, expire, IMAGE_MODE_RESIZE_TO_CROP, size...)
	originalPath := util.NewImagePath(uri, dir, expire, IMAGE_MODE_ORIGINAL)

	// When the image exists, return the image path
	if !util.IsEmptyImage(path) {
		return C.CString(strings.Replace(path, dir, "", -1))
	}

	inputImg, imgPath, isImageExist := operate.InputImage(uri, originalPath)
	if !isImageExist {
		return C.CString(uri)
	}
	outputImg := operate.ResizeToCrop(imgPath, size, inputImg)
	_, filePath, _ := operate.CreateImage(outputImg, path)

	fileName := strings.Replace(filePath, dir, "", -1)
	return C.CString(fileName)
}
