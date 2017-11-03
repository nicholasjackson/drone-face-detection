package gocv

/*
#include <stdlib.h>
#include "core.h"
*/
import "C"
import (
	"image"
	"unsafe"
)

// MatType is the type for the various different kinds of Mat you can create.
type MatType int

const (
	// MatTypeCV8U is a Mat of 8-bit unsigned int
	MatTypeCV8U MatType = 0

	// MatTypeCV8S is a Mat of 8-bit signed int
	MatTypeCV8S = 1

	// MatTypeCV16U is a Mat of 16-bit unsigned int
	MatTypeCV16U = 2

	// MatTypeCV16S is a Mat of 16-bit signed int
	MatTypeCV16S = 3

	// MatTypeCV32S is a Mat of 32-bit signed int
	MatTypeCV32S = 4

	// MatTypeCV32F is a Mat of 32-bit float
	MatTypeCV32F = 5

	// MatTypeCV64F is a Mat of 64-bit float
	MatTypeCV64F = 6
)

// Mat represents an n-dimensional dense numerical single-channel
// or multi-channel array. It can be used to store real or complex-valued
// vectors and matrices, grayscale or color images, voxel volumes,
// vector fields, point clouds, tensors, and histograms.
//
// For further details, please see:
// http://docs.opencv.org/3.3.1/d3/d63/classcv_1_1Mat.html
//
type Mat struct {
	p C.Mat
}

// NewMat returns a new empty Mat.
func NewMat() Mat {
	return Mat{p: C.Mat_New()}
}

// NewMatWithSize returns a new Mat with a specific size and type.
func NewMatWithSize(rows int, cols int, mt MatType) Mat {
	return Mat{p: C.Mat_NewWithSize(C.int(rows), C.int(cols), C.int(mt))}
}

// NewMatFromScalar returns a new Mat for a specific Scalar value
func NewMatFromScalar(s Scalar, mt MatType) Mat {
	sVal := C.struct_Scalar{
		val1: C.double(s.Val1),
		val2: C.double(s.Val2),
		val3: C.double(s.Val3),
		val4: C.double(s.Val4),
	}

	return Mat{p: C.Mat_NewFromScalar(sVal, C.int(mt))}
}

// Close the Mat object.
func (m *Mat) Close() error {
	C.Mat_Close(m.p)
	m.p = nil
	return nil
}

// Ptr returns the Mat's underlying object pointer.
func (m *Mat) Ptr() C.Mat {
	return m.p
}

// Empty determines if the Mat is empty or not.
func (m *Mat) Empty() bool {
	isEmpty := C.Mat_Empty(m.p)
	return isEmpty != 0
}

// Clone returns a cloned full copy of the Mat.
func (m *Mat) Clone() Mat {
	return Mat{p: C.Mat_Clone(m.p)}
}

// CopyTo copies Mat into destination Mat.
func (m *Mat) CopyTo(dst Mat) {
	C.Mat_CopyTo(m.p, dst.p)
	return
}

// Region returns a new Mat that points to a region of this Mat. Changes made to the
// region Mat will affect the original Mat, since they are pointers to the underlying
// OpenCV Mat object.
func (m *Mat) Region(rio image.Rectangle) Mat {
	cRect := C.struct_Rect{
		x:      C.int(rio.Min.X),
		y:      C.int(rio.Min.Y),
		width:  C.int(rio.Size().X),
		height: C.int(rio.Size().Y),
	}

	return Mat{p: C.Mat_Region(m.p, cRect)}
}

// Rows returns the number of rows for this Mat.
func (m *Mat) Rows() int {
	return int(C.Mat_Rows(m.p))
}

// Cols returns the number of columns for this Mat.
func (m *Mat) Cols() int {
	return int(C.Mat_Cols(m.p))
}

// GetUCharAt returns a value from a specific row/col in this Mat expecting it to
// be of type uchar aka CV_8U.
func (m *Mat) GetUCharAt(row int, col int) int8 {
	return int8(C.Mat_GetUChar(m.p, C.int(row), C.int(col)))
}

// GetSCharAt returns a value from a specific row/col in this Mat expecting it to
// be of type schar aka CV_8S.
func (m *Mat) GetSCharAt(row int, col int) int8 {
	return int8(C.Mat_GetSChar(m.p, C.int(row), C.int(col)))
}

// GetShortAt returns a value from a specific row/col in this Mat expecting it to
// be of type short aka CV_16S.
func (m *Mat) GetShortAt(row int, col int) int16 {
	return int16(C.Mat_GetShort(m.p, C.int(row), C.int(col)))
}

// GetIntAt returns a value from a specific row/col in this Mat expecting it to
// be of type int aka CV_32S.
func (m *Mat) GetIntAt(row int, col int) int32 {
	return int32(C.Mat_GetInt(m.p, C.int(row), C.int(col)))
}

// GetFloatAt returns a value from a specific row/col in this Mat expecting it to
// be of type float aka CV_32F.
func (m *Mat) GetFloatAt(row int, col int) float32 {
	return float32(C.Mat_GetFloat(m.p, C.int(row), C.int(col)))
}

// GetDoubleAt returns a value from a specific row/col in this Mat expecting it to
// be of type double aka CV_64F.
func (m *Mat) GetDoubleAt(row int, col int) float64 {
	return float64(C.Mat_GetDouble(m.p, C.int(row), C.int(col)))
}

// AbsDiff calculates the per-element absolute difference between two arrays
// or between an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#ga6fef31bc8c4071cbc114a758a2b79c14
//
func AbsDiff(src1 Mat, src2 Mat, dst Mat) {
	C.Mat_AbsDiff(src1.p, src2.p, dst.p)
}

// Add calculates the per-element sum of two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#ga10ac1bfb180e2cfda1701d06c24fdbd6
//
func Add(src1 Mat, src2 Mat, dst Mat) {
	C.Mat_Add(src1.p, src2.p, dst.p)
}

// AddWeighted calculates the weighted sum of two arrays.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#gafafb2513349db3bcff51f54ee5592a19
//
func AddWeighted(src1 Mat, alpha float64, src2 Mat, beta float64, gamma float64, dst Mat) {
	C.Mat_AddWeighted(src1.p, C.double(alpha),
		src2.p, C.double(beta), C.double(gamma), dst.p)
}

// BitwiseAnd computes bitwise conjunction of the two arrays (dst = src1 & src2).
// Calculates the per-element bit-wise conjunction of two arrays
// or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#ga60b4d04b251ba5eb1392c34425497e14
//
func BitwiseAnd(src1 Mat, src2 Mat, dst Mat) {
	C.Mat_BitwiseAnd(src1.p, src2.p, dst.p)
}

// BitwiseNot inverts every bit of an array.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#ga0002cf8b418479f4cb49a75442baee2f
//
func BitwiseNot(src1 Mat, dst Mat) {
	C.Mat_BitwiseNot(src1.p, dst.p)
}

// BitwiseOr calculates the per-element bit-wise disjunction of two arrays
// or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#gab85523db362a4e26ff0c703793a719b4
//
func BitwiseOr(src1 Mat, src2 Mat, dst Mat) {
	C.Mat_BitwiseOr(src1.p, src2.p, dst.p)
}

// BitwiseXor calculates the per-element bit-wise "exclusive or" operation
// on two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#ga84b2d8188ce506593dcc3f8cd00e8e2c
//
func BitwiseXor(src1 Mat, src2 Mat, dst Mat) {
	C.Mat_BitwiseXor(src1.p, src2.p, dst.p)
}

// InRange checks if array elements lie between the elements of two Mat arrays.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#ga48af0ab51e36436c5d04340e036ce981
//
func InRange(src Mat, lb Mat, ub Mat, dst Mat) {
	C.Mat_InRange(src.p, lb.p, ub.p, dst.p)
}

// GetOptimalDFTSize returns the optimal Discrete Fourier Transform (DFT) size
// for a given vector size.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#ga6577a2e59968936ae02eb2edde5de299
//
func GetOptimalDFTSize(vecsize int) int {
	return int(C.Mat_GetOptimalDFTSize(C.int(vecsize)))
}

// DFT performs a forward or inverse Discrete Fourier Transform (DFT)
// of a 1D or 2D floating-point array.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#gadd6cf9baf2b8b704a11b5f04aaf4f39d
//
func DFT(src Mat, dst Mat) {
	C.Mat_DFT(src.p, dst.p)
}

// Merge creates one multi-channel array out of several single-channel ones.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#ga7d7b4d6c6ee504b30a20b1680029c7b4
//
func Merge(src Mat, count int, dst Mat) {
	C.Mat_Merge(src.p, C.size_t(count), dst.p)
}

// NormType for normalization operations.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#gad12cefbcb5291cf958a85b4b67b6149f
//
type NormType int

const (
	NormInf      NormType = 1
	NormL1                = 2
	NormL2                = 4
	NormL2Sqr             = 5
	NormHamming           = 6
	NormHamming2          = 7
	NormTypeMask          = 7
	NormRelative          = 8
	NormMixMax            = 32
)

// Normalize normalizes the norm or value range of an array.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d2/de8/group__core__array.html#ga87eef7ee3970f86906d69a92cbf064bd
//
func Normalize(src Mat, dst Mat, alpha float64, beta float64, typ NormType) {
	C.Mat_Normalize(src.p, dst.p, C.double(alpha), C.double(beta), C.int(typ))
}

// Scalar is a 4-element vector widely used in OpenCV to pass pixel values.
//
// For further details, please see:
// http://docs.opencv.org/3.3.1/d1/da0/classcv_1_1Scalar__.html
//
type Scalar struct {
	Val1 float64
	Val2 float64
	Val3 float64
	Val4 float64
}

// NewScalar returns a new Scalar. These are usually colors typically being in BGR order.
func NewScalar(v1 float64, v2 float64, v3 float64, v4 float64) Scalar {
	s := Scalar{Val1: v1, Val2: v2, Val3: v3, Val4: v4}
	return s
}

func toGoBytes(b C.struct_ByteArray) []byte {
	return C.GoBytes(unsafe.Pointer(b.data), b.length)
}
