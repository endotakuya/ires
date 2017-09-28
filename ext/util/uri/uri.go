package util

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"io/ioutil"
)

// ローカルの画像かどうか
func IsLocalFile(path string) bool {
	if strings.Index(path, "http") == -1 {
		return true
	} else {
		return false
	}
}

// 画像名を取得
func getImageName(uri string) (string, string) {
	splitPath 	:= strings.Split(uri, "/")
	fileName 	:= splitPath[len(splitPath)-1]		// ex. sample.jpg
	ext 		:= filepath.Ext(fileName)			// ex. .jpg
	return strings.Replace(fileName, ext, "", 1), ext
}

// 画像のフルパスを生成
func NewImagePath(uri, dir, expire string, mode int, size ...int) string {
	name := NewImageName(uri, expire)
	return FilePath(name, dir, mode, size...)
}

// ファイル名を生成
func NewImageName(uri, expire string) string {
	imageName, ext := getImageName(uri)
	fullImageName := expire + "_" + imageName + ext
	return fullImageName
}

// 画像を保存するパスの設定
func FilePath(name string, d string, mode int, size ...int) string {
	paths := []rune(d)
	pathsLastIndex := len(paths) - 1
	lastChar := string(paths[pathsLastIndex])
	if lastChar == "/" {
		d = string(paths[:pathsLastIndex])
	}

	// 画像格納先
	var oDir string
	switch mode {
	case 0: oDir = "ires/resize"
	case 1: oDir = "ires/crop"
	case 2: oDir = "ires/resize_to_crop"
	case 3: oDir = "ires/original"
	}

	var prefix string
	if len(size) == 2 {
		prefix = PrefixSize(size...)
	} else {
		prefix = "original"
	}
	oDir += "/" + prefix

	// Create directory
	oPath := filepath.Join(d, oDir)
	if _, err := os.Stat(oPath); err != nil {
		if err := os.MkdirAll(oPath, 0777); err != nil {
			panic(err)
		}
	}

	pullPath := filepath.Join(oPath, name)
	return pullPath
}

// サイズからプレフィックスを作成 ex. 640x480
func PrefixSize(size ...int) string {
	prefix := strconv.Itoa(size[0]) + "x" + strconv.Itoa(size[1])
	return prefix
}

// リサイズ済みのファイルがあれば、処理せず返す
func IsEmptyImage(path string) bool {
	_, err := ioutil.ReadFile(path)
	if err != nil {
		return true
	} else {
		return false
	}
}