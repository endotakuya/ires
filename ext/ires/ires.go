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

// Resize is ...
func (i *Ires) Resize() string {
	i.Mode = Resize
	distURI := i.imageURI(false)
	// When the image exists, return the image path
	if isExistsImage(distURI) {
		return i.targetImageURI(distURI)
	}

	if !i.inputImage() {
		return i.URI
	}

	var outputImg image.Image
	if i.validResizeType() {
		outputImg = resize.Resize(uint(i.Width), uint(i.Height), i.InputImage.Image, resize.Lanczos3)
	} else {
		outputImg = i.InputImage.Image
	}
	createImage(outputImg, distURI, i.InputImage.Format)
	return i.targetImageURI(distURI)
}

// Crop is Crop ...
func (i *Ires) Crop() string {
	i.Mode = Crop
	distURI := i.imageURI(false)
	// When the image exists, return the image path
	if isExistsImage(distURI) {
		return i.targetImageURI(distURI)
	}

	if !i.inputImage() {
		return i.URI
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
	return i.targetImageURI(distURI)
}

// ResizeToCrop is ...
func (i *Ires) ResizeToCrop() string {
	i.Mode = ResizeToCrop
	distURI := i.imageURI(false)
	// When the image exists, return the image path
	if isExistsImage(distURI) {
		return i.targetImageURI(distURI)
	}

	if !i.inputImage() {
		return i.URI
	}

	var outputImg image.Image
	if i.validResizeType() {
		outputImg = i.resizeToCrop()
	} else {
		outputImg = i.InputImage.Image
	}
	createImage(outputImg, distURI, i.InputImage.Format)
	return i.targetImageURI(distURI)
}
