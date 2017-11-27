package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/nats-io/nats"
	messages "github.com/nicholasjackson/drone-messages"
)

var processing = false
var faceProcessor *FaceProcessor
var nc *nats.Conn
var natsServer = flag.String("nats", "nats://localhost:4222", "connection string for nats server")

func main() {
	flag.Parse()

	var err error
	nc, err = nats.Connect(*natsServer)
	if err != nil {
		log.Fatal("Unable to connect to nats")
	}

	faceProcessor = NewFaceProcessor()

	sub, _ := nc.Subscribe(messages.MessageDroneImage, func(m *nats.Msg) {
		if !processing {
			go processMessage(m)
		}
	})
	defer sub.Unsubscribe()

	startServer()

	handleExit()
}

func handleExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func processMessage(m *nats.Msg) {
	processing = true
	defer func() { processing = false }()

	filename := "./latest.jpg"
	di := messages.DroneImage{}
	di.DecodeMessage(m.Data)
	di.SaveDataToFile(filename)

	faces, bounds := faceProcessor.DetectFaces(filename)
	if len(faces) > 0 {
		fdm := messages.FaceDetected{
			Faces:  faces,
			Bounds: bounds,
		}

		nc.Publish(messages.MessageFaceDetection, fdm.EncodeMessage())
	}
}

func startServer() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":4000", nil)
}
