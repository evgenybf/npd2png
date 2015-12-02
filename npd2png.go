package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	width  = 320
	height = 190
)

func replaceExt(path string, ext string) string {
	return path[0:len(path)-len(filepath.Ext(path))] + ext
}

func convert(inputFile string, outputFile string) error {
	fileIn, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer fileIn.Close()

	fileIn.Seek(14, 0)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	buffer := make([]byte, width*3)

	for row := 0; row < height; row++ {
		n, err := io.ReadFull(fileIn, buffer)
		if err != nil {
			return err
		}
		if n < width {
			break
		}
		col := 0
		for i := 0; i < len(buffer); i += 3 {
			c := color.RGBA{buffer[i], buffer[i+1], buffer[i+2], 255}
			img.Set(col, row, c)
			col++
		}
	}

	fileOut, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	err = png.Encode(fileOut, img)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	files, err := filepath.Glob("*.NPD")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		outFile := replaceExt(file, ".png")
		log.Printf("Convert %s to %s", file, outFile)
		err := convert(file, outFile)
		if err != nil {
			log.Print("Error: ", err.Error())
		} else {
			os.Remove(file)
		}
	}
}
