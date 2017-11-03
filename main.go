package main

import (
	"fmt"
	"image"
	"image/color"
	"sort"

	"gocv.io/x/gocv"
)

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

func main() {
	deviceID := 1

	// open webcam
	webcam, err := gocv.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	tmpImg := gocv.NewMat()
	defer img.Close()
	defer tmpImg.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}
	green := color.RGBA{0, 255, 0, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	classifier.Load("./data/haarcascade_frontalface_default.xml")

	eyeclassifier := gocv.NewCascadeClassifier()
	defer eyeclassifier.Close()
	eyeclassifier.Load("./data/haarcascade_eye.xml")

	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}

		if img.Empty() {
			continue
		}

		gocv.Resize(img, img, image.Point{}, 0.5, 0.5, gocv.InterpolationNearestNeighbor)
		gocv.CvtColor(img, tmpImg, gocv.ColorRGBToGray)

		// detect faces
		faces := classifier.DetectMultiScale(tmpImg)
		fmt.Printf("found %d faces\n", len(faces))

		if len(faces) > 0 {
			sort.Sort(BySize(faces))

			// draw a rectangle around each face on the original image
			gocv.Rectangle(img, faces[0], blue, 3)

			eyeImg := tmpImg.Region(faces[0])
			eyes := eyeclassifier.DetectMultiScale(eyeImg)
			for _, r := range eyes {
				gocv.Rectangle(img, r, green, 3)
			}
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		gocv.WaitKey(1)
	}
}
