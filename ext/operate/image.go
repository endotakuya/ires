package operate

import (
	"image"
	"image/jpeg"
	"net/http"
	"os"

	"github.com/endotakuya/ires/ext/util/uri"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

// 入力画像
func InputImage(uri, path string) (image.Image, string, bool) {
	if util.IsLocalFile(uri) {
		return LocalImage(uri), uri, true
	} else {
		img, path, isImageExist := DownloadImage(uri, path)
		return img, path, isImageExist
	}
}

// http経由での画像を保存
func DownloadImage(uri, path string) (image.Image, string, bool) {
	res, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, path, false
	}
	return CreateImage(img, path)
}

// 画像を作成
func CreateImage(img image.Image, path string) (image.Image, string, bool) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jpeg.Encode(file, img, nil)

	return LocalImage(path), path, true
}

// ローカルの画像を取得
func LocalImage(uri string) image.Image {
	file, err := os.Open(uri)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	// Decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}
	return img
}

// リサイズ + 切り取り
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

// 変更するサイズと画像のアスペクト比が一致するかどうか
func isValidAspectRatio(path string, size []int) (bool, image.Config) {
	conf := imageConfig(path)
	aspH := (conf.Height * size[0]) / conf.Width
	if aspH == size[1] {
		return true, conf
	} else {
		return false, conf
	}
}

// 入力画像の情報を取得
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

// アスペクト比の異なる画像に対するモード設定
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