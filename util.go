package framereader

import (
	"github.com/lazywei/go-opencv/opencv"
	"unsafe"
)

func DecodeImageMem(data []byte) *opencv.IplImage {
	buf := opencv.CreateMatHeader(1, len(data), opencv.CV_8U)
	buf.SetData(unsafe.Pointer(&data[0]), opencv.CV_AUTOSTEP)
	defer buf.Release()

	return opencv.DecodeImage(unsafe.Pointer(buf), opencv.CV_LOAD_IMAGE_UNCHANGED)
}
