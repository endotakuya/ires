package ires

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/goware/urlx"
)

// Generate image path
func (i *Ires) imageURI(original bool) string {
	paths := []rune(i.Dir)
	pathsLastIndex := len(paths) - 1
	lastChar := string(paths[pathsLastIndex])
	dir := i.Dir
	if lastChar == "/" {
		dir = string(paths[:pathsLastIndex])
	}

	var oDir string
	if i.IsLocal {
		oDir = localPath(i.Mode)
	} else {
		oDir = remotePath(i.URI)
	}

	// Create directory
	oPath := filepath.Join(dir, oDir)
	if _, err := os.Stat(oPath); err != nil {
		if err := os.MkdirAll(oPath, 0777); err != nil {
			panic(err)
		}
	}

	name := i.imageName(original)
	return filepath.Join(oPath, name)
}

// Generate image name
func (i *Ires) imageName(original bool) string {
	splitPath := strings.Split(i.URI, "/")

	// ex. sample.jpg
	fileName := splitPath[len(splitPath)-1]
	// ex. .jpg
	ext := filepath.Ext(fileName)

	name := strings.Replace(fileName, ext, "", 1)

	extInfo := strings.Split(ext, "?")
	if len(extInfo) > 1 {
		ext = extInfo[0]
		name += "_" + strings.Join(extInfo[1:], "")
	}

	var prefix string
	if original {
		prefix = "original"
	} else {
		switch i.Mode {
		case Resize:
			prefix = prefixSize(i.Size) + "_resize"
		case Crop:
			prefix = prefixSize(i.Size) + "_crop"
		case ResizeToCrop:
			prefix = prefixSize(i.Size) + "_resize_to_crop"
		}
	}

	return i.Expire + "_" + name + "_" + prefix + ext
}

// Create prefix by size
// ex. 640x480
func prefixSize(s Size) string {
	return strconv.Itoa(s.Width) + "x" + strconv.Itoa(s.Height)
}

// リサイズ済みのファイルがあれば、処理せず返す
func isExistsImage(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// Read directory
func (i *Ires) readImageDir() string {
	var dir string
	if i.IsLocal {
		dir = localPath(i.Mode)
	} else {
		dir = remotePath(i.URI)
	}
	return filepath.Join(i.Dir, dir)
}

// if local image, create ires directory
func localPath(mode Mode) string {
	var dir string
	switch mode {
	case Resize:
		dir = "ires/resize"
	case Crop:
		dir = "ires/crop"
	case ResizeToCrop:
		dir = "ires/resize_to_crop"
	}
	return dir
}

// if http image, parse URL & make directory
func remotePath(uri string) string {
	u, err := urlx.Parse(uri)
	dir := []string{"ires"}
	if err != nil {
		panic(err)
	}

	dir = append(dir, u.Host)
	path := strings.Split(u.Path, "/")
	dir = append(dir, path[1:len(path)-1]...)

	return strings.Join(dir, "/")
}

// Optimize image path
func (i *Ires) targetImageURI(uri string) string {
	return strings.Replace(uri, i.Dir, "", -1)
}
