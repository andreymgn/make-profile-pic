package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
)

var (
	red   = color.RGBA{0xff, 0, 0, 0xff}
	green = color.RGBA{0, 0xff, 0, 0xff}
	blue  = color.RGBA{0, 0, 0xff, 0xff}
	white = color.RGBA{0xfa, 0xfa, 0xfa, 0xff}
	// AllColors is an array of all available colors for squares in generated image
	AllColors = []color.RGBA{red, green, blue}
)

// Avatar contains data about avatar to be generated
type Avatar struct {
	SideInPixels  int
	SideInSquares int
	Color         color.RGBA
	Squares       [][]bool
}

// NewAvatar returns new Avatar instance with parameters Color and Squares unitialized
func NewAvatar(sideInPixels, sideInSquares int) *Avatar {
	return &Avatar{SideInPixels: sideInPixels,
		SideInSquares: sideInSquares,
		Color:         white,
		Squares:       nil,
	}
}

// Generate fills out random data in Avatar like squares positions and squares color
func (a *Avatar) Generate(filename string) {
	img := make([][]bool, a.SideInSquares)
	for i := range img {
		img[i] = make([]bool, a.SideInSquares)
	}
	for y := 0; y < a.SideInSquares; y++ {
		for x := 0; x < a.SideInSquares/2+1; x++ {
			img[x][y] = rand.Int()%2 == 0
			img[a.SideInSquares-x-1][y] = img[x][y]
		}
	}
	a.Squares = img
	colorIndex := int(rand.Int31()) % len(AllColors)
	a.Color = AllColors[colorIndex]
	a.draw(filename)
}

func (a *Avatar) draw(filename string) {
	upperLeft := image.Point{0, 0}
	lowerRight := image.Point{a.SideInPixels, a.SideInPixels}
	img := image.NewRGBA(image.Rectangle{upperLeft, lowerRight})
	squareSide := a.SideInPixels / a.SideInSquares
	for y := 0; y < a.SideInSquares; y++ {
		for x := 0; x < a.SideInSquares; x++ {
			imgY := y * squareSide
			imgX := x * squareSide
			point := image.Point{imgX, imgY}
			square := makeSquare(point, squareSide)
			color := white
			if a.Squares[x][y] {
				color = a.Color
			}
			draw.Draw(img, square, &image.Uniform{color}, image.ZP, draw.Src)
		}
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("error while creating file: ", err)
	}
	png.Encode(f, img)
}

func makeSquare(upperLeft image.Point, side int) image.Rectangle {
	lowerRight := image.Point{upperLeft.X + side, upperLeft.Y + side}
	return image.Rectangle{upperLeft, lowerRight}
}
