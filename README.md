# go-framereader
A simple API to Iterate over video frames sent over STDIN using MJPEG or PPM formats. The resulting frames are OpenCV Image objects.

## Example
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

### Running

##### MJPEG
```bash
ffmpeg -i VIDEO.mp4 -qscale:v 2 -vcodec mjpeg -f image2pipe -pix_fmt yuvj420p - | ./example
```
##### PPM

```bash
ffmpeg -i VIDEO.mp4 -f image2pipe -qscale:v 2 -vcodec ppm | ./example
```
