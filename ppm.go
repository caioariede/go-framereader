package framereader

import (
	"bufio"
	"github.com/lazywei/go-opencv/opencv"
	"os"
)

func PPMStdinReader(bufsize int) (ch chan *opencv.IplImage) {
	ch = make(chan *opencv.IplImage)

	go func() {
		data := []byte{}
		reader := bufio.NewReader(os.Stdin)
		first := true

		for {
			bytes := make([]byte, bufsize)

			if numBytes, err := reader.Read(bytes); err != nil {
				panic(err)
			} else {
				for i := 0; i < numBytes; i++ {
					magic := false
					if bytes[i] == 'P' {
						if (numBytes-1) > i && bytes[i+1] == '6' {
							magic = true
						}
					}
					if magic {
						if !first {
							ch <- DecodeImageMem(data)
							data = []byte{}
						} else {
							first = false
						}
					}

					data = append(data, bytes[i])
				}
			}
		}
	}()

	return
}
