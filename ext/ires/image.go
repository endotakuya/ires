package ires

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)


// Input image
func inputImage(i *Ires) (image.Image, string, bool) {
	if isLocalFile(i.Uri) {
		i.IsLocal = true
		img, format := localImage(i.Uri)
		return img, format, true
	} else {
		return downloadImage(i)
	}
}


// Save http image
func downloadImage(i *Ires) (image.Image, string, bool) {
	res, err := http.Get(i.Uri)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	header, r := copyReader(res.Body)
	format := formatSearch(r)

	img, _, err := image.Decode(io.MultiReader(header, res.Body))
	if err != nil {
		return nil, "", false
	}
	return createImage(img, i.imagePath(IMAGE_MODE_ORIGINAL), format)
}


func createImage(img image.Image, path, format string) (image.Image, string, bool) {
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

	return img, format, true
}

// Load image
func localImage(uri string) (image.Image, string) {
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


// Resizing & Cropping
func resizeToCrop(i *Ires, inputImg image.Image) image.Image {
	var outputImg image.Image
	path := i.imagePath(IMAGE_MODE_ORIGINAL)
	isAsp, conf := isValidAspectRatio(path, i.Size)

	width  := i.Size.Width
	height := i.Size.Height

	if isAsp {
		outputImg = resize.Resize(uint(width), uint(height), inputImg, resize.Lanczos3)
	} else {
		var resizeImg image.Image

		// Resize
		mode := resizeMode(conf, i.Size)
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


// Verify aspect ratio
func isValidAspectRatio(path string, s Size) (bool, image.Config) {
	conf := imageConfig(path)
	aspH := (conf.Height * s.Width) / conf.Width
	if aspH == s.Height {
		return true, conf
	} else {
		return false, conf
	}
}


// Image config
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


// Select image resize mode
func resizeMode(conf image.Config, s Size) int {
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
func formatSearch(r io.Reader) string{
	_, format, err := image.DecodeConfig(r)
	if err != nil {
		return "jpeg"
	}
	fmt.Println(format)
	return format
}


// Copy Reader
func copyReader(body io.Reader) (io.Reader, io.Reader) {
	header := bytes.NewBuffer(nil)
	r := io.TeeReader(body, header)
	return header, r
}