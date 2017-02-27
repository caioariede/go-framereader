package framereader

import (
	"github.com/lazywei/go-opencv/opencv"
)

func VideoFileReader(filename string) (ch chan *opencv.IplImage) {
	ch = make(chan *opencv.IplImage)

	go func() {
		cap := opencv.NewFileCapture(filename)

		if cap == nil {
			panic("cannot open video")
		}

		defer cap.Release()

		fps := int(cap.GetProperty(opencv.CV_CAP_PROP_FPS))

		for {
			img := cap.QueryFrame()

			if img == nil {
				break
			}

			ch <- img

			opencv.WaitKey(1000 / fps)
		}
	}()

	return
}
