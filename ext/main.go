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
)

func init() {}
func main() {}

//export resizeImage
func resizeImage(Uri *C.char, width, height int, Dir *C.char) *C.char {
	dir := C.GoString(Dir)
	uri := C.GoString(Uri)

	inputImg, _ := operate.InputImage(uri, dir)
	outputImg 	:= resize.Resize(uint(width), uint(height), inputImg, resize.Lanczos3)

	size := []int{width, height}
	path := util.NewImagePath(uri, dir, IMAGE_MODE_RESIZE, size...)
	_, filePath := operate.CreateImage(outputImg, path)

	fileName := strings.Replace(filePath, dir, "", -1)
	return C.CString(fileName)
}

//export cropImage
func cropImage(Uri *C.char, width, height int, Dir *C.char) *C.char {
	dir := C.GoString(Dir)
	uri := C.GoString(Uri)

	inputImg, _ := operate.InputImage(uri, dir)
	outputImg, _ := cutter.Crop(inputImg, cutter.Config{
		Width:  width,
		Height: height,
		Mode: cutter.Centered,
		Options: cutter.Copy,
	})

	size := []int{width, height}
	path := util.NewImagePath(uri, dir, IMAGE_MODE_CROP, size...)
	_, filePath := operate.CreateImage(outputImg, path)

	fileName := strings.Replace(filePath, dir, "", -1)
	return C.CString(fileName)
}

//export resizeToCropImage
func resizeToCropImage(Uri *C.char, width, height int, Dir *C.char) *C.char {
	dir := C.GoString(Dir)
	uri := C.GoString(Uri)

	inputImg, imgPath := operate.InputImage(uri, dir)
	size := []int{width, height}
	outputImg := operate.ResizeToCrop(imgPath, size, inputImg)

	path := util.NewImagePath(uri, dir, IMAGE_MODE_RESIZE_TO_CROP, size...)
	_, filePath := operate.CreateImage(outputImg, path)

	fileName := strings.Replace(filePath, dir, "", -1)
	return C.CString(fileName)
}
