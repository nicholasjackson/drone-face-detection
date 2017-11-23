package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"sort"

	"gocv.io/x/gocv"
)

// BySize allows sorting images by size
type BySize []image.Rectangle

func (s BySize) Len() int {
	return len(s)
}
func (s BySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s BySize) Less(i, j int) bool {
	return s[i].Size().X > s[j].Size().X && s[i].Size().Y > s[j].Size().Y
}

var blue = color.RGBA{0, 0, 255, 0}

type FaceProcessor struct {
	classifier *gocv.CascadeClassifier
}

func NewFaceProcessor() *FaceProcessor {
	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	classifier.Load("./data/haarcascade_frontalface_default.xml")

	return &FaceProcessor{&classifier}
}

func (fp *FaceProcessor) DetectFaces(file string) {
	log.Println("Detect faces")

	img := gocv.IMRead(file, gocv.IMReadColor)
	defer img.Close()

	gocv.CvtColor(img, img, gocv.ColorRGBToGray)
	gocv.Resize(img, img, image.Point{}, 0.3, 0.3, gocv.InterpolationArea)

	// detect faces
	faces := fp.classifier.DetectMultiScaleWithParams(img, 1.07, 8, 0, image.Point{X: 30, Y: 30}, image.Point{X: 100, Y: 100})

	fmt.Printf("found %d faces\n", len(faces))

	if len(faces) > 0 {
		sort.Sort(BySize(faces))

		// draw a rectangle around each face on the original image
		for _, f := range faces {
			gocv.Rectangle(img, f, blue, 1)
		}

		gocv.IMWrite("./detect.jpg", img)
		log.Println("Written image\n")
	}
}
