package main

import (
	"C"

	"github.com/endotakuya/ires/ext/ires"
)

func init() {}
func main() {}

//export iresImagePath
func iresImagePath(Uri *C.char, width, height int, Mode, Dir, Expire *C.char) *C.char {
	uri		:= C.GoString(Uri)
	mode 	:= C.GoString(Mode)
	dir 	:= C.GoString(Dir)
	expire	:= C.GoString(Expire)

	r := &ires.Ires{
		Uri: uri,
		Size: ires.Size{
			Width: width,
			Height: height,
		},
		Dir: dir,
		Expire: expire,
		IsLocal: false,
	}

	var imagePath string
	switch mode {
	case "resize":
		imagePath = r.Resize()
	case "crop":
		imagePath = r.Crop()
	case "resize_to_crop":
		imagePath = r.ResizeToCrop()
	}

	return C.CString(imagePath)
}