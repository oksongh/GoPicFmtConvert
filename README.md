# GoPicFmtConvert

* 画像のフォーマットを変換するGo製のCUIツールです。
* 現在のところjpg,png,gifに対応しています。
* それ、ffmpegでできるよ

## usage
* -f:フォーマット
* -o:出力先ディレクトリ(指定しなかったら元のファイルと同じ場所に出力)

```bash
GoPicFmtConvert -f jpg -o test_out testcase/*.png 
```

## example
pngのファイルをjpgに
```bash
GoPicFmtConvert -f jpg -o test_out testcase/*.png 
```

移動
```bash
mkdir dir
mv *.jpg dir
```

bzipで圧縮
```
tar -cvzf dir.tar.gz dir/
```

gzipで圧縮
```bash
tar -cvjf dir.tar.bz2 dir/
```
