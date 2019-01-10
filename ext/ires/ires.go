package ires

import (
	"image"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

type Mode int

const (
	Resize Mode = iota
	Crop
	ResizeToCrop
)

type ResizeType int

const (
	All ResizeType = iota
	Smaller
	Larger
)

type (
	// Size is ...
	Size struct {
		Width, Height int
	}

	// InputImage is ...
	InputImage struct {
		Image  image.Image
		Config image.Config
		Format string
		URI    string
	}

	// Ires is ...
	Ires struct {
		Size
		ResizeType
		Mode
		*InputImage
		URI     string
		Dir     string
		Expire  string
		IsLocal bool
	}
)

// Init is ...
func Init(size Size, rType int, uri, dir, expire string) (i *Ires) {
	i = &Ires{
		Size: 		size,
		ResizeType: ResizeType(rType),
		URI:        uri,
		Dir:        dir,
		Expire:     expire,
	}

	// If local image, True
	i.CheckLocal()

	// Delete the expiration date image
	i.DeleteExpireImage()

	return
}

// Resize is ...
func (i *Ires) Resize() (string, error) {
	i.Mode = Resize
	distURI, err := i.imageURI(false)
	if err != nil {
		return "", err
	}
	// When the image exists, return the image path
	if isExistsImage(distURI) {
		return i.targetImageURI(distURI), nil
	}

	if err := i.inputImage(); err != nil {
		return "", err
	}

	var outputImg image.Image
	if i.validResizeType() {
		outputImg = resize.Resize(uint(i.Width), uint(i.Height), i.InputImage.Image, resize.Lanczos3)
	} else {
		outputImg = i.InputImage.Image
	}
	createImage(outputImg, distURI, i.InputImage.Format)
	return i.targetImageURI(distURI), nil
}

// Crop is Crop ...
func (i *Ires) Crop() (string, error) {
	i.Mode = Crop
	distURI, err := i.imageURI(false)
	if err != nil {
		return "", err
	}
	// When the image exists, return the image path
	if isExistsImage(distURI) {
		return i.targetImageURI(distURI), nil
	}

	if err := i.inputImage(); err != nil {
		return "", err
	}

	var outputImg image.Image
	if i.validResizeType() {
		outputImg, _ = cutter.Crop(i.InputImage.Image, cutter.Config{
			Width:   i.Width,
			Height:  i.Height,
			Mode:    cutter.Centered,
			Options: cutter.Copy,
		})
	} else {
		outputImg = i.InputImage.Image
	}
	createImage(outputImg, distURI, i.InputImage.Format)
	return i.targetImageURI(distURI), nil
}

// ResizeToCrop is ...
func (i *Ires) ResizeToCrop() (string, error) {
	i.Mode = ResizeToCrop
	distURI, err := i.imageURI(false)
	if err != nil {
		return "", err
	}
	// When the image exists, return the image path
	if isExistsImage(distURI) {
		return i.targetImageURI(distURI), nil
	}

	if err := i.inputImage(); err != nil {
		return "", err
	}

	var outputImg image.Image
	if i.validResizeType() {
		if img, err := i.resizeToCrop(); err != nil {
			return "", err
		} else {
			outputImg = img
		}
	} else {
		outputImg = i.InputImage.Image
	}
	createImage(outputImg, distURI, i.InputImage.Format)
	return i.targetImageURI(distURI), nil
}
