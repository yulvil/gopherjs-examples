package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randomColor() color.NRGBA {
	return color.NRGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
}

func base64img(width, height, blocks int) string {
	var cellsizex int = width / blocks
	var cellsizey int = height / blocks

	colors := make([][]color.NRGBA, blocks)
	for i, _ := range colors {
		colors[i] = make([]color.NRGBA, blocks)
		for j := 0; j < blocks; j++ {
			colors[i][j] = randomColor()
		}
	}

	// 32 x 4  (2)
	// cellsizex 16
	// cellsizey 2
	// color[16][2]

	m := image.NewNRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{width, height}})
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var xx int = x / cellsizex
			var yy int = y / cellsizey
			// fmt.Println(x, y, xx, yy)
			var c = colors[xx][yy]
			m.SetNRGBA(x, y, c)
		}
	}

	buf := bytes.NewBuffer([]byte{})
	if err := png.Encode(buf, m); err != nil {
		fmt.Println(err)
	}
	b64str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return b64str
}

func main() {

	if js.Global != nil {
		// running from a browser
		js.Global.Set("gopkg", map[string]interface{}{
			"base64img": base64img,
		})
	} else {
		width := 32
		height := 4
		blocks := 4
		b64str := base64img(width, height, blocks)
		fmt.Println(b64str)
	}

}
