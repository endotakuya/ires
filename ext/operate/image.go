package operate

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"

	"github.com/endotakuya/ires/ext/util/uri"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)


func InputImage(uri, path string) (image.Image, string, bool) {
	if util.IsLocalFile(uri) {
		img, format := LocalImage(uri)
		return img, format, true
	} else {
		img, path, isImageExist := DownloadImage(uri, path)
		return img, path, isImageExist
	}
}


// Save http image
func DownloadImage(uri, path string) (image.Image, string, bool) {
	res, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	header, r := copyReader(res.Body)
	format := formatSearch(r)

	img, _, err := image.Decode(io.MultiReader(header, res.Body))
	if err != nil {
		return nil, path, false
	}
	return CreateImage(img, path, format)
}


func CreateImage(img image.Image, path, format string) (image.Image, string, bool) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	switch format {
	case "jpeg":
		jpeg.Encode(file, img, nil)
	case "png":
		png.Encode(file, img)
	case "gif":
		gif.Encode(file, img, nil)
	default:
		jpeg.Encode(file, img, nil)
	}

	return img, path, true
}


func LocalImage(uri string) (image.Image, string) {
	file, err := os.Open(uri)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	// Decode jpeg into image.Image
	header, r := copyReader(file)
	format := formatSearch(r)

	img, _, err := image.Decode(io.MultiReader(header, file))
	if err != nil {
		panic(err)
	}
	return img, format
}


func ResizeToCrop(path string, size []int, inputImg image.Image) image.Image {
	var outputImg image.Image
	isAsp, conf := isValidAspectRatio(path, size)

	if isAsp {
		outputImg = resize.Resize(uint(size[0]), uint(size[1]), inputImg, resize.Lanczos3)
	} else {
		var resizeImg image.Image

		// Resize
		mode := resizeMode(conf, size)
		switch mode {
		case 1, 3:
			resizeImg = resize.Resize(uint(size[0]), 0, inputImg, resize.Lanczos3)
		case 2, 4:
			resizeImg = resize.Resize(0, uint(size[1]), inputImg, resize.Lanczos3)
		default:
			resizeImg = inputImg
		}

		// Cropping
		outputImg, _ = cutter.Crop(resizeImg, cutter.Config{
			Width:  size[0],
			Height: size[1],
			Mode: cutter.Centered,
			Options: cutter.Copy,
		})

	}
	return outputImg
}


func isValidAspectRatio(path string, size []int) (bool, image.Config) {
	conf := imageConfig(path)
	aspH := (conf.Height * size[0]) / conf.Width
	if aspH == size[1] {
		return true, conf
	} else {
		return false, conf
	}
}


func imageConfig(path string) image.Config {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	conf, _, err := image.DecodeConfig(file)
	if err != nil {
		panic(err)
	}
	return conf
}


func resizeMode(conf image.Config, size []int) int {
	if conf.Width >= conf.Height && size[0] >= size[1] {
		return 1
	} else if conf.Width >= conf.Height && size[0] < size[1] {
		return 2
	} else if conf.Width < conf.Height && size[0] >= size[1] {
		return 3
	} else if conf.Width < conf.Height && size[0] < size[1] {
		return 4
	}
	return 0
}


func formatSearch(r io.Reader) string{
	_, format, err := image.DecodeConfig(r)
	if err != nil {
		return "jpeg"
	}
	return format
}


func copyReader(body io.Reader) (io.Reader, io.Reader) {
	header := bytes.NewBuffer(nil)
	r := io.TeeReader(body, header)
	return header, r
}