package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/nats-io/nats"
	messages "github.com/nicholasjackson/drone-messages"
)

var processing = false
var faceProcessor *FaceProcessor

func main() {
	nc, err := nats.Connect("nats://192.168.1.113:4222")
	if err != nil {
		log.Fatal("Unable to connect to nats")
	}

	faceProcessor = NewFaceProcessor()

	sub, _ := nc.Subscribe(messages.MessageDroneImage, func(m *nats.Msg) {
		log.Println("New Message")

		if !processing {
			go processMessage(m)
		}
	})
	defer sub.Unsubscribe()

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

	faceProcessor.DetectFaces(filename)
}
