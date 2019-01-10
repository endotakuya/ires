package main

import (
	"C"
	"log"

	"github.com/endotakuya/ires/ext/ires"
)

func init() {}
func main() {}

//export resizeImagePath
func resizeImagePath(URI *C.char, width, height, rType int, Dir, Expire *C.char) *C.char {
	uri, dir, expire := cToString(URI, Dir, Expire)
	r := ires.Init(ires.Size{ Width: width, Height: height }, rType, uri, dir, expire)

	distURI, err := r.Resize()
	if err != nil {
		log.Print(err)
		return C.CString(r.URI)
	}
	return C.CString(distURI)
}

//export cropImagePath
func cropImagePath(URI *C.char, width, height, rType int, Dir, Expire *C.char) *C.char {
	uri, dir, expire := cToString(URI, Dir, Expire)
	r := ires.Init(ires.Size{ Width: width, Height: height }, rType, uri, dir, expire)

	distURI, err := r.Crop()
	if err != nil {
		log.Print(err)
		return C.CString(r.URI)
	}
	return C.CString(distURI)
}

//export resizeToCropImagePath
func resizeToCropImagePath(URI *C.char, width, height, rType int, Dir, Expire *C.char) *C.char {
	uri, dir, expire := cToString(URI, Dir, Expire)
	r := ires.Init(ires.Size{ Width: width, Height: height }, rType, uri, dir, expire)

	distURI, err := r.ResizeToCrop()
	if err != nil {
		log.Print(err)
		return C.CString(r.URI)
	}
	return C.CString(distURI)
}

// Convert *C.char to String
func cToString(uri, dir, expire *C.char) (u, d, e string){
	u = C.GoString(uri)
	d = C.GoString(dir)
	e = C.GoString(expire)
	return
}