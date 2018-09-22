package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"
)

func main() {
	f, err := os.Open("g.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	g, err := gif.DecodeAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(g.Image))
	fmt.Println(g.Config.Width, g.Config.Height)
	g1 := g.Image[59].SubImage(image.Rect(0, 0, g.Config.Width, g.Config.Height))

	file1, err := os.Create("test1.png")
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()

	if err != nil {
		fmt.Println(err)
	}
	png.Encode(file1, g1)
}
