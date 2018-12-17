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
	"strings"
	"time"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

// Check expiration date
func (i *Ires) DeleteExpireImage() error {
	today := time.Now().Format("20060102")
	dir, err := i.readImageDir()
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		findName := file.Name()
		matched, err := path.Match(today+"_*", findName)
		if err != nil {
			return err
		}
		if matched {
			if err := deleteImage(path.Join(dir, findName)); err != nil {
				return err
			}
		}
	}
	return nil
}

// Input image type is Local or HTTP
func (i *Ires) CheckLocal() {
	if strings.Index(i.URI, "http") == -1 {
		i.IsLocal = true
	} else {
		i.IsLocal = false
	}
}

// Input image
func (i *Ires) inputImage() error {
	if i.IsLocal {
		img, format := localImage(i.URI)
		i.InputImage = &InputImage{
			Image:  img,
			Format: format,
			URI:    i.URI,
		}
		i.setConfig()
		return nil
	}
	return i.downloadImage()
}

// Save http image
func (i *Ires) downloadImage() error {
	res, err := http.Get(i.URI)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, res.Body)
	format := formatSearch(res.Body)

	img, _, err := image.Decode(buf)
	if err != nil {
		return err
	}

	var distURI string
	if distURI, err = i.imageURI(true); err != nil {
		return err
	}

	i.InputImage = &InputImage{
		Image:  img,
		Format: format,
		URI:    distURI,
	}

	if createImage(img, distURI, format) != nil {
		return err
	} else {
		if err := i.setConfig(); err != nil {
			return err
		}
	}
	return nil
}

func createImage(img image.Image, path, format string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
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
	return nil
}

// Load image
func localImage(uri string) (image.Image, string) {
	file, err := os.Open(uri)
	if err != nil {
		return nil, ""
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, file)
	format := formatSearch(file)

	img, _, err := image.Decode(buf)
	if err != nil {
		return nil, ""
	}
	return img, format
}

// Set image config
func (i *Ires) setConfig() error {
	file, err := os.Open(i.InputImage.URI)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, file)

	conf, _, err := image.DecodeConfig(buf)
	if err != nil {
		return err
	}
	i.InputImage.Config = conf
	return nil
}

// Resizing & Cropping
func (i *Ires) resizeToCrop() (image.Image, error) {
	inputImg := i.InputImage.Image
	var outputImg image.Image
	isAsp, conf := i.isValidAspectRatio()

	width := i.Size.Width
	height := i.Size.Height

	if isAsp {
		outputImg = resize.Resize(uint(width), uint(height), inputImg, resize.Lanczos3)
	} else {
		var resizeImg image.Image

		// Resize
		mode := resizeMode(conf, i.Size)
		switch mode {
		case 3, 4:
			resizeImg = resize.Resize(uint(width), 0, inputImg, resize.Lanczos3)
		case 1, 2:
			resizeImg = resize.Resize(0, uint(height), inputImg, resize.Lanczos3)
		default:
			resizeImg = inputImg
		}

		// Cropping
		var err error
		if outputImg, err = cutter.Crop(resizeImg, cutter.Config{
			Width:   width,
			Height:  height,
			Mode:    cutter.Centered,
			Options: cutter.Copy,
		}); err != nil {
			return nil, err
		}
	}
	return outputImg, nil
}

// Delete image
func deleteImage(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	} else {
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	return nil
}

// Verify aspect ratio
func (i *Ires) isValidAspectRatio() (bool, image.Config) {
	conf := i.InputImage.Config
	s := i.Size
	println("\n")
	aspH := (conf.Height * s.Width) / conf.Width
	if aspH == s.Height {
		return true, conf
	}
	return false, conf
}

// Select image resize mode
func resizeMode(conf image.Config, s Size) int {
	srcWidth := s.Width
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
func formatSearch(r io.Reader) string {
	_, format, err := image.DecodeConfig(r)
	if err != nil {
		return "jpeg"
	}
	return format
}

// Valid resize type
func (i *Ires) validResizeType() bool {
	config := i.InputImage.Config
	valid := false
	switch i.ResizeType {
	case All:
		valid = true
	case Smaller:
		if config.Width < i.Width && config.Height < i.Height {
			valid = true
		}
	case Larger:
		if i.Width <= config.Width && i.Height <= config.Height {
			valid = true
		}
	default:
		valid = true
	}
	return valid
}
