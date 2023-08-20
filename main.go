package main

import (
	"fmt"
	"image"
	"image-concatenator/utils"
	_ "image/jpeg"
	png "image/png"
	"os"
)

const dir = "./"

func main() {
	fileNames, err := utils.FindFiles(dir)

	if err != nil {
		fmt.Println("Could not find files")
		panic(err)
	}

	images := utils.FindImages(fileNames)

	fmt.Println("Found", len(images), "images.")

	outImgHeight := utils.GetSumYOfImages(images)
	outImgWidth := utils.GetMaxXByImages(images)

	newImage := image.NewRGBA(image.Rect(0, 0, outImgWidth, outImgHeight))

	utils.AppendImages(images, newImage)

	out, err := os.Create("out.png")

	if err != nil {
		fmt.Println("Could not create file out.png")
		panic(err)
	}

	defer out.Close()

	fmt.Println("Encoding image...")

	err = png.Encode(out, newImage)

	if err != nil {
		fmt.Println("Could not encode image")
		panic(err)
	}

	fmt.Println("Done!")
}
