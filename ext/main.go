package main

import (
	"C"
	"log"

	"github.com/endotakuya/ires/ext/ires"
)

func init() {}
func main() {}

//export iresImagePath
func iresImagePath(URI *C.char, width, height, rType, mode int, Dir, Expire *C.char) *C.char {
	uri := C.GoString(URI)
	dir := C.GoString(Dir)
	expire := C.GoString(Expire)

	r := &ires.Ires{
		Size: ires.Size{
			Width:  width,
			Height: height,
		},
		ResizeType: ires.ResizeType(rType),
		URI:        uri,
		Dir:        dir,
		Expire:     expire,
	}

	// If local image, True
	r.CheckLocal()
	// Delete the expiration date image
	r.DeleteExpireImage()

	var distURI string
	var err error
	switch ires.Mode(mode) {
	case ires.Resize:
		distURI, err = r.Resize()
	case ires.Crop:
		distURI, err = r.Crop()
	case ires.ResizeToCrop:
		distURI, err = r.ResizeToCrop()
	}

	if err != nil {
		log.Print(err)
		return C.CString(r.URI)
	}
	return C.CString(distURI)
}
