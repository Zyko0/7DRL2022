package assets

import (
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomonobold"
)

var (
	ScoreFontFace         font.Face
	CardTitleFontFace     font.Face
	CardBodyTitleFontFace font.Face
	CardBodyTextFontFace  font.Face
)

func init() {
	pfont, err := truetype.Parse(gomonobold.TTF)
	if err != nil {
		log.Fatal(err)
	}
	ScoreFontFace = truetype.NewFace(pfont, &truetype.Options{
		Size: 36,
	})

	CardTitleFontFace = truetype.NewFace(pfont, &truetype.Options{
		Size: 24,
	})

	CardBodyTitleFontFace = truetype.NewFace(pfont, &truetype.Options{
		Size: 16,
	})

	CardBodyTextFontFace = truetype.NewFace(pfont, &truetype.Options{
		Size: 12,
	})
}
