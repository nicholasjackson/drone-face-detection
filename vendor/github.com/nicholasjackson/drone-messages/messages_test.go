package messages

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

var flightData = []byte{59, 255, 129, 3, 1, 1, 6, 70, 108, 105, 103, 104, 116, 1, 255, 130, 0, 1, 3, 1, 11, 66, 97, 115, 101, 77, 101, 115, 115, 97, 103, 101, 1, 255, 132, 0, 1, 7, 67, 111, 109, 109, 97, 110, 100, 1, 12, 0, 1, 5, 86, 97, 108, 117, 101, 1, 4, 0, 0, 0, 23, 255, 131, 3, 1, 1, 11, 66, 97, 115, 101, 77, 101, 115, 115, 97, 103, 101, 1, 255, 132, 0, 0, 0, 11, 255, 130, 1, 0, 1, 4, 84, 101, 115, 116, 0}

func TestEncodesFlightMessage(t *testing.T) {
	is := is.New(t)

	m := Flight{Command: "Test"}
	data := m.EncodeMessage()

	is.True(len(data) > 1) // data length should be greater than 0
	fmt.Println(data)
}

func TestDecodesFlightMessage(t *testing.T) {
	is := is.New(t)

	m := Flight{}
	m.DecodeMessage(flightData)

	is.Equal("Test", m.Command) // decoded command should be test
}
