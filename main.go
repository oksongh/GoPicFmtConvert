package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

// https://qiita.com/KemoKemo/items/d135ddc93e6f87008521
func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func getFullPathWithoutExt(path string) string {
	return path[:len(path)-len(filepath.Ext(path))]
}
func exitOnError(err error) {

	if err != nil {
		panic(err)
		// fmt.Println(err.Error())
		// os.Exit(-1)
	}
}

// func find(tgt string, src []string) bool {
// 	for _, v := range src {
// 		if tgt == v {
// 			return true
// 		}
// 	}
// 	return false
// }

type Encode func(io.Writer, image.Image) error

var jpegEncode Encode = func(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}
var gifEncode Encode = func(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}

var fmt2Encoder map[string]Encode = map[string]Encode{
	"jpeg": jpegEncode,
	"jpg":  jpegEncode,
	"png":  png.Encode,
	"gif":  gifEncode,
}

func convert(in, out, outFmt string) {
	// ディレクトリ除外
	if finfo, err := os.Stat(in); err != nil || finfo.IsDir() {
		return
	}

	fin, err := os.Open(in)
	exitOnError(err)
	defer fin.Close()

	// 画像以外除外,元の画像の形式
	image, _, err := image.Decode(fin)
	if err != nil {
		return
	}

	base := getFullPathWithoutExt(in)

	fo, err := os.Create(base + "." + outFmt)
	exitOnError(err)
	defer fo.Close()

	if encode, ok := fmt2Encoder[outFmt]; !ok {
		exitOnError(errors.New("not supported image format"))
	} else {
		encode(fo, image)
	}
}

// go run main.go filename fileextension
func main() {

	// var output string
	// flag.StringVar(&output, "o", "", "output directry")
	flag.Parse()

	srcGlob := flag.Arg(0)
	targetFmt := flag.Arg(1)

	if srcGlob == "" || targetFmt == "" {
		exitOnError(errors.New("invalid args error"))
	}

	src, err := filepath.Glob(srcGlob)
	exitOnError(err)

	fmt.Println(src)

	for _, fname := range src {

		// ディレクトリ除外
		if finfo, err := os.Stat(fname); err != nil || finfo.IsDir() {
			continue
		}

		fin, err := os.Open(fname)
		exitOnError(err)
		defer fin.Close()

		// 画像以外除外,元の画像の形式
		image, _, err := image.Decode(fin)
		if err != nil {
			continue
		}

		base := getFullPathWithoutExt(fname)

		fo, err := os.Create(base + "." + targetFmt)
		exitOnError(err)
		defer fo.Close()

		if encode, ok := fmt2Encoder[targetFmt]; !ok {
			exitOnError(errors.New("not supported image format"))
		} else {
			encode(fo, image)
		}
	}

}
