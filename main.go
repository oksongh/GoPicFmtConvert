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
	"strings"
)

// https://qiita.com/KemoKemo/items/d135ddc93e6f87008521
// hoge/fuga/baz.txt => baz
func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// hoge/fuga/baz.txt => hoge/fuga/baz
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

func convert(inPath, outPath, outFmt string) error {

	// ディレクトリ除外
	if finfo, err := os.Stat(inPath); err != nil {
		return err
	} else if finfo.IsDir() {
		return fmt.Errorf("%s is directory", inPath)
	}

	fin, err := os.Open(inPath)

	if err != nil {
		return err
	}
	defer fin.Close()

	// 入力の形式に対応してるか
	encode, ok := fmt2Encoder[outFmt]
	if !ok {
		return fmt.Errorf("%s is not a supported image format", outFmt)
	}

	// 画像以外のファイルを除外,変換前の画像
	image, _, err := image.Decode(fin)
	if err != nil {
		return err
	}

	fo, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer fo.Close()

	encode(fo, image)
	return nil
}

// go run main.go filename fileextension
func main() {

	var uiOutdir, outFmt string
	flag.StringVar(&uiOutdir, "o", "", "output directry")

	fmts := make([]string, 0, len(fmt2Encoder))
	for key := range fmt2Encoder {
		fmts = append(fmts, key)
	}
	flag.StringVar(&outFmt, "f", "", "format "+strings.Join(fmts, "|"))
	flag.Parse()

	src := flag.Args()

	if len(src) == 0 || outFmt == "" {
		exitOnError(errors.New("invalid args error"))
	}

	fmt.Println(src)

	// 入力の形式に対応してるか
	if _, ok := fmt2Encoder[outFmt]; !ok {
		exitOnError(fmt.Errorf("%s is not a supported image format", outFmt))
	}

	for _, fin := range src {

		outdir := ""
		if uiOutdir == "" {
			outdir = filepath.Dir(fin)
		} else {
			outdir = uiOutdir
		}

		base := getFileNameWithoutExt(fin)
		fo := filepath.Join(outdir, base) + "." + outFmt

		err := convert(fin, fo, outFmt)
		if err != nil {
			fmt.Println(err)
		}

	}

}
