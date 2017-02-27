# go-framereader
A simple API to Iterate over video frames sent over STDIN using MJPEG or PPM formats. The resulting frames are OpenCV Image objects.

## example
```go
import (
    "github.com/caioariede/go-framereader"
    "github.com/lazywei/go-opencv/opencv"
)

func main() {
    win := opencv.NewWindow("GoOpenCV: VideoPlayer")
    defer win.Destroy()

    // Read MJPEG stream from STDIN in chunks of 10240 bytes
    reader := framereader.MJPEGStdinReader(10240)

    for image := range reader {
        win.ShowImage(image)
        opencv.WaitKey(1000)
    }
}
```
