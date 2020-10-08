package main

import (
	"fmt"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

var inputPicture string = "Test.jpg"
var outputPicture string = "saveFile.jpg"

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 0 {
		inputPicture = argsWithoutProg[0]
		outputPicture = argsWithoutProg[1]
		fmt.Println("Input Full Path", inputPicture)
		fmt.Println("Output Full Path", outputPicture)
	} else {
		fmt.Println("Using Default")

	}

	img := gocv.IMRead(inputPicture, 1)
	defer img.Close()

	harrcascade := "haarcascade_frontalface_default.xml"
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(harrcascade)
	defer classifier.Close()

	color := color.RGBA{0, 255, 0, 0}

	rects := classifier.DetectMultiScale(img)
	for _, r := range rects {
		fmt.Println("detected", r)
		gocv.Rectangle(&img, r, color, 3)
	}

	gocv.IMWrite(outputPicture, img)

}
