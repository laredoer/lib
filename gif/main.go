package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io/ioutil"
	"os"
	"reflect"
	"time"

	"github.com/nfnt/resize"
)

var asciiChar = "MND8OZ$7I?+=~:,.."

func main() {
	//打开要转换的gif
	f, err := os.Open("g.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//gif解码
	g, err := gif.DecodeAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(g.Image))
	fmt.Println(g.Config.Width, g.Config.Height)
	//获取每一帧图片数据
	images := Create(g.Image, g.Config.Width, g.Config.Height)
	//创建临时文件夹
	PathExists("test")
	if err != nil {
		panic(err)
	}

	for _, v := range images {
		// file1, err := os.Create(fmt.Sprintf("test/%d.png", k))
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// png.Encode(file1, v)
		// file1.Close()

		finalASCIIArt := asciiArt(getHeight(v, 120))
		fmt.Println(string(finalASCIIArt))
		time.Sleep(time.Millisecond * 100)
	}

}

// func CreatePaletted(length int) []*image.Paletted {

// 	f2, err := os.Open("123.png")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer f2.Close()
// 	g2, err := png.Decode(f2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	p2 := image.NewPaletted(image.Rect(0, 0, 200, 200), palette.Plan9)
// 	draw.Draw(p2, p2.Bounds(), g2, image.ZP, draw.Src) //添加图片
// }

func asciiArt(img image.Image, w, h int) []byte {
	table := []byte(asciiChar)
	buffer := new(bytes.Buffer)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buffer.WriteByte(table[pos])
		}
		_ = buffer.WriteByte('\n')
	}
	return buffer.Bytes()
}

func getHeight(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	height := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(height), img, resize.Lanczos3)
	return img, w, height
}

func Create(p []*image.Paletted, w, h int) (i []image.Image) {
	for _, v := range p {
		i = append(i, v.SubImage(image.Rect(0, 0, w, h)))
	}
	return i
}

func PathExists(path string) {
	_, err := os.Stat(path)
	if err == nil {

	}
	if os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname + fi.Name() + "\\")
		} else {
			fmt.Println(fi.Name())
		}
	}
	return err
}
