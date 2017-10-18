package ires

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)


// Input image
func InputImage(i *Ires) (image.Image, string, bool) {
	if i.IsLocal {
		img, format := LocalImage(i.Uri)
		return img, format, true
	} else {
		return DownloadImage(i)
	}
}


// Save http image
func DownloadImage(i *Ires) (image.Image, string, bool) {
	res, err := http.Get(i.Uri)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	header, r := CopyReader(res.Body)
	format := FormatSearch(r)

	img, _, err := image.Decode(io.MultiReader(header, res.Body))
	if err != nil {
		return nil, "", false
	}
	return CreateImage(img, i.ImagePath(IMAGE_MODE_ORIGINAL), format), format, true
}


func CreateImage(img image.Image, path, format string) image.Image {
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
	return img
}


// Load image
func LocalImage(uri string) (image.Image, string) {
	file, err := os.Open(uri)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	// Decode jpeg into image.Image
	header, r := CopyReader(file)
	format := FormatSearch(r)

	img, _, err := image.Decode(io.MultiReader(header, file))
	if err != nil {
		panic(err)
	}
	return img, format
}


// Resizing & Cropping
func ResizeToCrop(i *Ires, inputImg image.Image) image.Image {
	var outputImg image.Image
	path := i.ImagePath(IMAGE_MODE_ORIGINAL)
	isAsp, conf := IsValidAspectRatio(path, i.Size)

	width  := i.Size.Width
	height := i.Size.Height

	if isAsp {
		outputImg = resize.Resize(uint(width), uint(height), inputImg, resize.Lanczos3)
	} else {
		var resizeImg image.Image

		// Resize
		mode := ResizeMode(conf, i.Size)
		switch mode {
		case 1, 3:
			resizeImg = resize.Resize(uint(width), 0, inputImg, resize.Lanczos3)
		case 2, 4:
			resizeImg = resize.Resize(0, uint(height), inputImg, resize.Lanczos3)
		default:
			resizeImg = inputImg
		}

		// Cropping
		outputImg, _ = cutter.Crop(resizeImg, cutter.Config{
			Width:  width,
			Height: height,
			Mode: cutter.Centered,
			Options: cutter.Copy,
		})

	}
	return outputImg
}


// Check expiration date
func (i *Ires) DeleteExpireImage(mode int) {
	today := time.Now().Format("20060102")
	dir := i.ReadImageDir(mode)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range files {
		findName := file.Name()
		matched,_ := path.Match(today + "_*", findName)
		if matched {
			DeleteImage(path.Join(dir, findName))
		}
	}
}


// Delete image
func DeleteImage(path string) {
	_, err := os.Stat(path)
	if err == nil {
		if err := os.Remove(path); err != nil {
			panic(err)
		}
	}
}


// Verify aspect ratio
func IsValidAspectRatio(path string, s Size) (bool, image.Config) {
	conf := ImageConfig(path)
	aspH := (conf.Height * s.Width) / conf.Width
	if aspH == s.Height {
		return true, conf
	} else {
		return false, conf
	}
}


// Image config
func ImageConfig(path string) image.Config {
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


// Select image resize mode
func ResizeMode(conf image.Config, s Size) int {
	srcWidth  := s.Width
	srcHeight := s.Height

	if conf.Width >= conf.Height && srcWidth >= srcHeight {
		return 1
	} else if conf.Width >= conf.Height && srcWidth < srcHeight {
		return 2
	} else if conf.Width < conf.Height && srcWidth >= srcHeight {
		return 3
	} else if conf.Width < conf.Height && srcWidth < srcHeight {
		return 4
	}
	return 0
}


// Search image format
// if defined, return "jpeg"
func FormatSearch(r io.Reader) string{
	_, format, err := image.DecodeConfig(r)
	if err != nil {
		return "jpeg"
	}
	return format
}


// Copy Reader
func CopyReader(body io.Reader) (io.Reader, io.Reader) {
	header := bytes.NewBuffer(nil)
	r := io.TeeReader(body, header)
	return header, r
}