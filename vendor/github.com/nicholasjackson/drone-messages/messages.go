package messages

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"image"
	"io"
	"io/ioutil"
	"os"
)

const (
	// MessageFlight is the name of a flight message for the drone
	MessageFlight = "drone.flight"
	// MessageFaceDetection is the name of a message when a new face has been detected
	MessageFaceDetection = "image.facedetection"
	// MessageDroneImage is the name of a messeage when the drone takes a new image
	MessageDroneImage = "image.new"
)

const (
	// CommandTakeOff instructs a drone taking off
	CommandTakeOff = "takeoff"
	// CommandLand instructs a drone to land
	CommandLand = "land"
)

// DroneImage defines a new image taken from a drone
type DroneImage struct {
	Data []byte // Gzipped data
}

// Flight defines a flight instruction message
type Flight struct {
	Command string
	Value   int
}

// FaceDetected defines a face detection message
type FaceDetected struct {
	Faces  []image.Rectangle
	Bounds image.Rectangle
}

// EncodeMessage gob encodes the message and returns a byte slice
func (bm *Flight) EncodeMessage() []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(bm)

	return b.Bytes()
}

// DecodeMessage decodes the messgage from gob byte slice
func (bm *Flight) DecodeMessage(data []byte) {
	gob.NewDecoder(bytes.NewBuffer(data)).Decode(bm)
}

// EncodeMessage gob encodes the message and returns a byte slice
func (bm *FaceDetected) EncodeMessage() []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(bm)

	return b.Bytes()
}

// DecodeMessage decodes the messgage from gob byte slice
func (bm *FaceDetected) DecodeMessage(data []byte) {
	gob.NewDecoder(bytes.NewBuffer(data)).Decode(bm)
}

// EncodeMessage gob encodes the message and returns a byte slice
func (bm *DroneImage) EncodeMessage() []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(bm)

	return b.Bytes()
}

// DecodeMessage decodes the messgage from gob byte slice
func (bm *DroneImage) DecodeMessage(data []byte) {
	gob.NewDecoder(bytes.NewBuffer(data)).Decode(bm)
}

// UnzippedData returns the message data in unzipped format
func (bm *DroneImage) UnzippedData() []byte {
	zr, _ := gzip.NewReader(bytes.NewBuffer(bm.Data))
	d, _ := ioutil.ReadAll(zr)
	return d
}

// SetZippedData gzips the given data and sets the Data field
func (bm *DroneImage) SetZippedData(raw []byte) {
	var zb bytes.Buffer
	zw, _ := gzip.NewWriterLevel(&zb, gzip.BestCompression)
	zw.Write(raw)
	zw.Close()

	bm.Data = zb.Bytes()
}

// SaveDataToFile uncompresses the Data field and saves to a file
func (bm *DroneImage) SaveDataToFile(filename string) error {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		os.Remove(filename)
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	d := bm.UnzippedData()

	io.Copy(f, bytes.NewBuffer(d))

	return nil
}
