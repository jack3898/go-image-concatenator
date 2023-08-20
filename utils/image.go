package utils

import (
	"fmt"
	"image"
	"os"

	"github.com/disintegration/imaging"
)

// Returns a slice of images from the given paths
// If an image could not be decoded, it will be skipped
func FindImages(paths []string) []image.Image {
	var files []image.Image

	for _, path := range paths {
		file, err := os.Open(path)

		if err != nil {
			fmt.Println("Could not open file", file)
			panic(err)
		}

		defer file.Close()

		image, _, imageErr := image.Decode(file)

		if imageErr != nil {
			fmt.Println("Could not decode", file.Name(), "skipping")

			continue
		}

		files = append(files, image)
	}

	return files
}

// Appends images to the out image vertically
func AppendImages(images []image.Image, out *image.RGBA) *image.RGBA {
	for i, img := range images {
		fmt.Println("Processing image", i+1, "of", len(images))

		var y = SliceReduce(images[:i], func(acc int, cur image.Image) int {
			return acc + cur.Bounds().Max.Y
		})

		for x := 0; x < img.Bounds().Max.X; x++ {
			for y2 := 0; y2 < img.Bounds().Max.Y; y2++ {
				out.Set(x, y+y2, img.At(x, y2))
			}
		}
	}

	return out
}

// Returns the sum of all Y values of all images
func GetSumYOfImages(images []image.Image) int {
	return SliceReduce(images, func(acc int, cur image.Image) int {
		return acc + cur.Bounds().Max.Y
	})
}

// Returns the maximum X value of all images
func GetMaxXByImages(images []image.Image) int {
	return SliceReduce(images, func(acc int, cur image.Image) int {
		if cur.Bounds().Max.X > acc {
			return cur.Bounds().Max.X
		}

		return acc
	})
}

// Scales an image by the given width, maintains aspect ratio
func ScaleImageByWidth(img image.Image, width int) image.Image {
	return imaging.Resize(img, width, 0, imaging.Lanczos)
}
