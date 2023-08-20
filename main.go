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
		panic("Could not find files")
	}

	images := utils.FindImages(fileNames)
	outImgWidth := utils.GetMaxXByImages(images)

	fmt.Println("Found", len(images), "images.")

	resizedImages := utils.SliceMap(images, func(image image.Image, index int) image.Image {
		if image.Bounds().Max.X == outImgWidth {
			fmt.Println("Image", index+1, "is already the correct size")
			return image
		}

		fmt.Println("Resizing image", index+1, "of", len(images))
		return utils.ScaleImageByWidth(image, outImgWidth)
	})

	// Calculate the height of the output image using the newly resized images
	outImgHeight := utils.GetSumYOfImages(resizedImages)

	outImage := image.NewRGBA(image.Rect(0, 0, outImgWidth, outImgHeight))

	utils.AppendImages(resizedImages, outImage)

	out, err := os.Create("out.png")

	if err != nil {
		panic("Could not create file out.png")
	}

	defer out.Close()

	fmt.Println("Encoding image")

	err = png.Encode(out, outImage)

	if err != nil {
		panic("Could not encode image")
	}

	fmt.Println("Done!")
}
