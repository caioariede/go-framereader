package framereader

import (
	"bufio"
	"github.com/lazywei/go-opencv/opencv"
	"os"
)

func MJPEGStdinReader(bufsize int) (ch chan *opencv.IplImage) {
	ch = make(chan *opencv.IplImage)

	go func() {
		data := []byte{}

		reader := bufio.NewReader(os.Stdin)

		ff := false
		skip := true
		imgReady := false

		for {
			bytes := make([]byte, bufsize)
			if numBytes, err := reader.Read(bytes); err != nil {
				panic(err)
			} else {
				for i := 0; i < numBytes; i++ {
					c := bytes[i]
					if ff && c == 0xd8 {
						skip = false
						data = append(data, 0xff)
					} else if ff && c == 0xd9 {
						imgReady = true
						data = append(data, c)
					}

					ff = c == 0xff

					if !skip {
						data = append(data, c)
					}

					if imgReady {
						ch <- DecodeImageMem(data)
						data = []byte{}
						imgReady = false
						skip = true
					}
				}
			}
		}
	}()

	return
}
