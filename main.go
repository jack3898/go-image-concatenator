package main

import (
	"fmt"
	"image"
	"image-concatenator/utils"
	_ "image/jpeg"
	png "image/png"
	"os"
)

func findImages() []string {
	entries, err := os.ReadDir("./images")

	if err != nil {
		fmt.Println("No directory 'images' found. Please create it, and put your images in numeric order.")
		panic(err)
	}

	files := utils.SliceFilter(entries, func(file os.DirEntry) bool {
		return !file.IsDir()
	})

	fileNames := utils.SliceMap(files, func(file os.DirEntry) string {
		return "./images/" + file.Name()
	})

	return fileNames
}

func readFiles(paths []string) []image.Image {
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
			fmt.Println("Could not decode image", file)

			continue
		}

		files = append(files, image)
	}

	return files
}

func main() {
	fileNames := findImages()
	images := readFiles(fileNames)

	fmt.Println("Found", len(images), "images.")

	var height = utils.SliceReduce(images, func(acc int, cur image.Image) int {
		return acc + cur.Bounds().Max.Y
	})

	var width = utils.SliceReduce(images, func(acc int, cur image.Image) int {
		if cur.Bounds().Max.X > acc {
			return cur.Bounds().Max.X
		}

		return acc
	})

	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	for i, img := range images {
		var y = utils.SliceReduce(images[:i], func(acc int, cur image.Image) int {
			return acc + cur.Bounds().Max.Y
		})

		for x := 0; x < img.Bounds().Max.X; x++ {
			for y2 := 0; y2 < img.Bounds().Max.Y; y2++ {
				newImage.Set(x, y+y2, img.At(x, y2))
			}
		}
	}

	out, err := os.Create("out.png")

	if err != nil {
		fmt.Println("Could not create file out.png")
		panic(err)
	}

	defer out.Close()

	err = png.Encode(out, newImage)

	if err != nil {
		fmt.Println("Could not encode image")
		panic(err)
	}

	fmt.Println("Done!")
}
