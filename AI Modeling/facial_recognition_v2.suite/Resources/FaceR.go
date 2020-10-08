package main

import (
	"fmt"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

var inputPicture string
var outputPicture string
var haarcascade string

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 0 {
		haarcascade = argsWithoutProg[0]
		inputPicture = argsWithoutProg[1]
		outputPicture = argsWithoutProg[2]
		fmt.Println("Input Full Path", inputPicture)
		fmt.Println("Output Full Path", outputPicture)
	} else {
		fmt.Println("No input or Output Found")
		fmt.Println("go run Facer.go  \"classifier.xml\" \"Input.png\" \"output.png\"")
		return

	}

	img := gocv.IMRead(inputPicture, 1)
	defer img.Close()

	//haarcascade := "haarcascade_frontalface_default.xml"
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(haarcascade)
	defer classifier.Close()

	color := color.RGBA{0, 255, 0, 0}

	rects := classifier.DetectMultiScale(img)
	for _, r := range rects {
		fmt.Println("detected", r)
		gocv.Rectangle(&img, r, color, 3)
	}

	gocv.IMWrite(outputPicture, img)

}
