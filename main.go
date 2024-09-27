package main

import (
	"image/png" // 注册 PNG 解码器
	"log"
	"os"
	"path/filepath"

	ico "github.com/Kodeworks/golang-image-ico"
	"github.com/disintegration/imaging"
)

func main() {
	inFullName := os.Args[1]
	dirName := filepath.Dir(inFullName)
	baseExtName := filepath.Base(inFullName)
	extName := filepath.Ext(baseExtName)
	if extName != ".png" {
		log.Panicln("input file must be png")
	}
	baseName := baseExtName[:len(baseExtName)-len(extName)]
	// 创建 favicon 输出文件
	outFullName := filepath.Join(dirName, baseName+".ico")
	outFull32Name := filepath.Join(dirName, baseName+"-32.ico")
	outFull64Name := filepath.Join(dirName, baseName+"-48.ico")

	inputFile, err := os.Open(inFullName)
	if err != nil {
		log.Printf("open input file failed: %v", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outFullName)
	if err != nil {
		log.Printf("create output file failed: %v", err)
		return
	}
	defer outputFile.Close()
	output32File, err := os.Create(outFull32Name)
	if err != nil {
		log.Printf("create output file failed: %v", err)
		return
	}
	defer output32File.Close()
	output48File, err := os.Create(outFull64Name)
	if err != nil {
		log.Printf("create output file failed: %v", err)
		return
	}
	defer output48File.Close()

	img, err := png.Decode(inputFile)
	if err != nil {
		log.Printf("decode input file failed: %v", err)
		return
	}
	// compress
	compressImg := imaging.Resize(img, 16, 16, imaging.Lanczos)
	compress32Img := imaging.Resize(img, 32, 32, imaging.Lanczos)
	compress48Img := imaging.Resize(img, 48, 48, imaging.Lanczos)
	err = ico.Encode(outputFile, compressImg)
	if err != nil {
		log.Printf("encode output file failed: %v", err)
		return
	}
	err = ico.Encode(output32File, compress32Img)
	if err != nil {
		log.Printf("encode output file failed: %v", err)
		return
	}
	err = ico.Encode(output48File, compress48Img)
	if err != nil {
		log.Printf("encode output file failed: %v", err)
		return
	}
	log.Printf("encode %s to %s successfully", inFullName, outFullName)
}
