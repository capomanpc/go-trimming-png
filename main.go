package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

// トリミングしたい範囲を指定
var (
	x0, y0        = 291, 93    // 開始座標 (左上)
	width, height = 1082, 1528 // トリミング後の幅と高さ
)

func main() {
	// トリミングするファイルがあるフォルダ
	// エスケープ処理
	folderPath := "C:\\Users\\capom\\Downloads\\20240305-0342_42136345a773d22975fa2416da561879\\sql\\"

	// フォルダ内の全てのファイルを取得
	files, err := filepath.Glob(folderPath + "/*.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		trimImage(file)
	}
}

func trimImage(filePath string) {
	// ファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// PNGをデコード
	img, err := png.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// トリミング範囲を計算
	rect := image.Rect(x0, y0, x0+width, y0+height)

	// トリミング実行
	trimmedImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(rect)

	outputDir := "C:\\Users\\capom\\Downloads\\20240305-0342_42136345a773d22975fa2416da561879\\trimmed\\"

	// トリミングした画像をPNGとして保存
	outputFilePath := outputDir + "trimmed_" + filepath.Base(filePath)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	png.Encode(outputFile, trimmedImg)
}
