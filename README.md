# GoPicFmtConvert

* 画像のフォーマットを変換するGo製のCUIツールです。
* 現在のところjpg,png,gifに対応しています。

## usage
* -f:フォーマット
* -o:出力先ディレクトリ(指定しなかったら元のファイルと同じ場所に出力)

```bash
GoPicFmtConvert -f jpg -o test_out testcase/*.png 
```

pngのファイルをjpgにした後、bzipで圧縮
```bash
GoPicFmtConvert -f jpg -o test_out testcase/*.png 
mkdir dir
tar -cvjf dir.tar.bz2 dir/
```
