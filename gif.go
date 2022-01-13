package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
)

func CreateGif(name string, slice []int, sorter func([]int, func(int, int) bool, func()), less func(int, int) bool) error {
	var palette = []color.Color{
		color.White,
		color.Black,
	}
	var frames []*image.Paletted
	var delays []int

	limit := len(slice)
	bounds := image.Rect(0, 0, limit, limit)

	sorter(slice, less, func() {
		frame := image.NewPaletted(bounds, palette)
		frames = append(frames, frame)
		delays = append(delays, 0)
		draw.Draw(frame, bounds, &image.Uniform{color.White}, image.ZP, draw.Src)
		for x := 0; x < limit; x++ {
			frame.Set(x, limit-slice[x], color.Black)
		}
	})

	f, err := os.OpenFile(name+".gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return gif.EncodeAll(f, &gif.GIF{
		Image: frames,
		Delay: delays,
	})
}
