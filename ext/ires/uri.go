package ires

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/goware/urlx"
)

// Input image type is Local or HTTP
func isLocalFile(path string) bool {
	if strings.Index(path, "http") == -1 {
		return true
	} else {
		return false
	}
}

// Generate image name
func imageName(i *Ires, mode int) string {
	splitPath 	:= strings.Split(i.Uri, "/")

	// ex. sample.jpg
	fileName := splitPath[len(splitPath)-1]
	// ex. .jpg
	ext := filepath.Ext(fileName)
	name := strings.Replace(fileName, ext, "", 1)
	ext = strings.Replace(ext, "?", "_", -1)

	var prefix string
	if mode == 3 {
		prefix = "original"
	} else {
		prefix = prefixSize(i.Size)
	}

	return i.Expire + "_" + name + "_" + prefix + ext
}

// Generate image path
func (i *Ires) imagePath(mode int) string {
	paths := []rune(i.Dir)
	pathsLastIndex := len(paths) - 1
	lastChar := string(paths[pathsLastIndex])
	dir := i.Dir
	if lastChar == "/" {
		dir = string(paths[:pathsLastIndex])
	}

	var oDir string
	if i.IsLocal {
		oDir = localPath(mode)
	} else {
		oDir = remotePath(i)
	}

	// Create directory
	oPath := filepath.Join(dir, oDir)
	if _, err := os.Stat(oPath); err != nil {
		if err := os.MkdirAll(oPath, 0777); err != nil {
			panic(err)
		}
	}

	name := imageName(i, mode)
	return filepath.Join(oPath, name)
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
	} else {
		return false
	}
}

// Delete image
//func (i *Ires) DeleteExpireImage(mode int) {
//
//	today := time.Now().Format("20060102")
//	path := i.ImagePath(mode) NewImagePath(uri, dir, today, mode, size...)
//	_, err := os.Stat(path)
//	if err == nil {
//		if err := os.Remove(path); err != nil {
//			panic(err)
//		}
//	}
//}


// if local image, create ires directory
func localPath(mode int) string {
	var dir string
	switch mode {
	case 0: dir = "ires/resize"
	case 1: dir = "ires/crop"
	case 2: dir = "ires/resize_to_crop"
	case 3: dir = "ires/original"
	}
	return dir
}


// if http image, parse URL & make directory
func remotePath(i *Ires) string {
	u, err := urlx.Parse(i.Uri)
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
func (i *Ires) targetImagePath(path string) string {
	return strings.Replace(path, i.Dir, "", -1)
}
