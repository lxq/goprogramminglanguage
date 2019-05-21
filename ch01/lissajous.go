// 产生随机生成的利萨茹图形GIF
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// slice的复合字面量
var palette = []color.Color{color.White, color.RGBA{0xFF, 0X00, 0X00, 0XFF}}

const (
	clrIdx1 = 0
	clrIdx2 = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8899", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5     //完整的X振荡器变化个数
		res    = 0.001 //角度分辨率
		size   = 100   //图像画面大小[-size .. +size]
		frames = 64    // 帧数
		delay  = 10    // ms延迟
	)

	freq := rand.Float64() * 3.0 // Y振荡器频率
	// struct的复合字面量
	anim := gif.GIF{LoopCount: frames}
	phase := 0.0
	for i := 0; i < frames; i++ {
		rc := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rc, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), clrIdx2)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
	// TODO: 未处理编码错误.
}
