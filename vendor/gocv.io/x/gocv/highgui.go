package gocv

/*
#include <stdlib.h>
#include "highgui.h"
*/
import "C"
import (
	"unsafe"
)

// Window is a wrapper around OpenCV's "HighGUI" named windows.
// While OpenCV was designed for use in full-scale applications and can be used
// within functionally rich UI frameworks (such as Qt*, WinForms*, or Cocoa*)
// or without any UI at all, sometimes there it is required to try functionality
// quickly and visualize the results. This is what the HighGUI module has been designed for.
//
// For further details, please see:
// http://docs.opencv.org/3.3.1/d7/dfc/group__highgui.html
//
type Window struct {
	name string
	open bool
}

// NewWindow creates a new named OpenCV window
//
// For further details, please see:
// http://docs.opencv.org/3.3.1/d7/dfc/group__highgui.html#ga5afdf8410934fd099df85c75b2e0888b
//
func NewWindow(name string) *Window {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.Window_New(cName, 1)

	return &Window{name: name, open: true}
}

// Close closes and deletes a named OpenCV Window.
//
// For further details, please see:
// http://docs.opencv.org/3.3.1/d7/dfc/group__highgui.html#ga851ccdd6961022d1d5b4c4f255dbab34
//
func (w *Window) Close() error {
	cName := C.CString(w.name)
	defer C.free(unsafe.Pointer(cName))

	C.Window_Close(cName)
	w.open = false
	return nil
}

// IsOpen checks to see if the Window seems to be open.
func (w *Window) IsOpen() bool {
	return w.open
}

// IMShow displays an image Mat in the specified window.
// This function should be followed by the WaitKey function which displays
// the image for specified milliseconds. Otherwise, it won't display the image.
//
// For further details, please see:
// http://docs.opencv.org/3.3.1/d7/dfc/group__highgui.html#ga453d42fe4cb60e5723281a89973ee563
//
func (w *Window) IMShow(img Mat) {
	cName := C.CString(w.name)
	defer C.free(unsafe.Pointer(cName))

	C.Window_IMShow(cName, img.p)
}

// WaitKey waits for a pressed key.
// This function is the only method in OpenCV's HighGUI that can fetch
// and handle events, so it needs to be called periodically
// for normal event processing
//
// For further details, please see:
// http://docs.opencv.org/3.3.1/d7/dfc/group__highgui.html#ga5628525ad33f52eab17feebcfba38bd7
//
func WaitKey(delay int) int {
	return int(C.Window_WaitKey(C.int(delay)))
}

// Trackbar is a wrapper around OpenCV's "HighGUI" window Trackbars.
type Trackbar struct {
	name   string
	parent *Window
}

// CreateTrackbar creates a trackbar and attaches it to the specified window.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d7/dfc/group__highgui.html#gaf78d2155d30b728fc413803745b67a9b
//
func (w *Window) CreateTrackbar(name string, max int) *Trackbar {
	cName := C.CString(w.name)
	defer C.free(unsafe.Pointer(cName))

	tName := C.CString(name)
	defer C.free(unsafe.Pointer(tName))

	C.Trackbar_Create(cName, tName, C.int(max))
	return &Trackbar{name: name, parent: w}
}

// GetPos returns the trackbar position.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d7/dfc/group__highgui.html#ga122632e9e91b9ec06943472c55d9cda8
//
func (t *Trackbar) GetPos() int {
	cName := C.CString(t.parent.name)
	defer C.free(unsafe.Pointer(cName))

	tName := C.CString(t.name)
	defer C.free(unsafe.Pointer(tName))

	return int(C.Trackbar_GetPos(cName, tName))
}

// SetPos sets the trackbar position.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d7/dfc/group__highgui.html#ga67d73c4c9430f13481fd58410d01bd8d
//
func (t *Trackbar) SetPos(pos int) {
	cName := C.CString(t.parent.name)
	defer C.free(unsafe.Pointer(cName))

	tName := C.CString(t.name)
	defer C.free(unsafe.Pointer(tName))

	C.Trackbar_SetPos(cName, tName, C.int(pos))
}

// SetMin sets the trackbar minimum position.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d7/dfc/group__highgui.html#gabe26ffe8d2b60cc678895595a581b7aa
//
func (t *Trackbar) SetMin(pos int) {
	cName := C.CString(t.parent.name)
	defer C.free(unsafe.Pointer(cName))

	tName := C.CString(t.name)
	defer C.free(unsafe.Pointer(tName))

	C.Trackbar_SetMin(cName, tName, C.int(pos))
}

// SetMax sets the trackbar maximum position.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d7/dfc/group__highgui.html#ga7e5437ccba37f1154b65210902fc4480
//
func (t *Trackbar) SetMax(pos int) {
	cName := C.CString(t.parent.name)
	defer C.free(unsafe.Pointer(cName))

	tName := C.CString(t.name)
	defer C.free(unsafe.Pointer(tName))

	C.Trackbar_SetMax(cName, tName, C.int(pos))
}
