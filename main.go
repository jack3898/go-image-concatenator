package main

import (
	"fmt"
	"image"
	"image-concatenator/utils"
	_ "image/jpeg"
	png "image/png"
	"math"
	"os"
	"regexp"
	"strconv"
)

const dir = "./"

func main() {
	fileNames, pngErr := utils.FindFiles(dir)

	if pngErr != nil {
		panic("Could not find files")
	}

	sortedFileNames := utils.SliceSort(fileNames, func(a, b string) int {
		re := regexp.MustCompile("[0-9]+")

		// convert file name to int
		aInt, aErr := strconv.Atoi(utils.SliceJoin(re.FindAllString(a, -1)))
		bInt, bErr := strconv.Atoi(utils.SliceJoin(re.FindAllString(b, -1)))

		if aErr != nil || bErr != nil {
			return 0
		}

		return aInt - bInt
	})

	images := utils.FindImages(sortedFileNames)
	imagesSortedByX := utils.SliceSort(images, func(imageA image.Image, imageB image.Image) int {
		return imageA.Bounds().Max.X - imageB.Bounds().Max.X
	})

	medianWidth := imagesSortedByX[len(imagesSortedByX)/2].Bounds().Max.X
	outImgWidth := int(math.Min(float64(medianWidth), 2000))

	fmt.Println("Found", len(images), "images.")
	fmt.Println("Median size of image collection is", medianWidth, "px.")

	if outImgWidth == 2000 {
		fmt.Println("NOTE: The output image is now clamped to a width of 2000px, to prevent an excessive file size.")
	}

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

	out, pngErr := os.Create("out.png")

	if pngErr != nil {
		panic("Could not create file out.png")
	}

	defer out.Close()

	fmt.Println("Encoding image")

	pngErr = png.Encode(out, outImage)

	if pngErr != nil {
		panic(pngErr)
	}

	fmt.Println("Done!")
}
