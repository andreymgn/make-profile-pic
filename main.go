package main

import (
	"flag"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	sizeInPixels := flag.Int("sizeInPixels", 200, "defines size of side in pixels")
	sizeInSquares := flag.Int("sizeInSquares", 5, "defines number of squares")
	filename := flag.String("f", "generated.png", "output file name")
	flag.Parse()
	a := NewAvatar(*sizeInPixels, *sizeInSquares)
	a.Generate(*filename)
}
