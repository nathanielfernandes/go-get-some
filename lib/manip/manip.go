package manip

import (
	"image"
	"main/lib/helpers"

	"github.com/fogleman/gg"
)

func StringStuff(title string, fields map[string]string) (image.Image, error) {
	const EMOJI_SIZE = 64
	const TITLE_SIZE = 80

	const W = 1128
	H := (len(fields)/3+1)*(EMOJI_SIZE+30) + TITLE_SIZE + 40

	dc := gg.NewContext(W, H)
	dc.SetHexColor("#bbddfb")
	dc.Clear()
	if err := dc.LoadFontFace("./fonts/coolvetica.ttf", TITLE_SIZE); err != nil {
		return nil, err
	}
	dc.SetHexColor("#37587d")
	dc.DrawString(title, 16, TITLE_SIZE)
	if err := dc.LoadFontFace("./fonts/coolvetica.ttf", 40); err != nil {
		return nil, err
	}
	dc.SetRGB(0, 0, 0)

	offSetY := 170
	offSetX := 30
	reset := 0
	for k, v := range fields {
		im, err := helpers.GetEmoji(k, EMOJI_SIZE)
		if err != nil {
			return nil, err
		}
		dc.DrawImageAnchored(im, offSetX, offSetY, 0.1, 0.5)
		dc.DrawStringAnchored(v, float64(offSetX+EMOJI_SIZE), float64(offSetY), 0, 0.5)

		if reset < 2 {
			offSetX += 370
			reset++
		} else {
			reset = 0
			offSetX = 30
			offSetY += EMOJI_SIZE + 30
		}
	}

	return dc.Image(), nil
}
