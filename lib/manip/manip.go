package manip

import (
	"bytes"
	"image"
	"main/lib/helpers"
	"strings"

	"github.com/fogleman/gg"
	qrcode "github.com/skip2/go-qrcode"
)

type dims struct {
	Name string
	X    float64
	Y    float64
}

var FIELD_POSITIONS = []dims{
	{"Name", 1026.0, 320.0},
	{"Pronouns", 1026.0, 455.0},
	{"Exp. Level", 1026.0, 590.0},
	{"Education", 620.0, 725.0},
	{"School", 620.0, 860.0},
	{"Fave Movie", 620.0, 995.0},
}

var DISCORD_PASS = helpers.GetImage("./static/pass.png")

// var BLANK = helpers.GetImage("./static/blank.png") // 360x360

const EOP = 1882.0

func DiscordPass(fields map[string]string) (image.Image, error) {
	const NAME_SIZE = 40
	const TEXT_COLOR = "#ffffff"

	const VALUE_SIZE = 50

	dc := gg.NewContextForImage(DISCORD_PASS)
	dc.SetHexColor(TEXT_COLOR)

	im, err := helpers.GetImageUrl(fields["Image"], 360, 360)
	if err == nil {
		pfpc := gg.NewContext(360, 360)
		pfpc.DrawRoundedRectangle(0, 0, 360.0, 360.0, 18.0)
		pfpc.Clip()
		pfpc.SetColor(image.White)

		color, ok := fields["color"]
		if !ok {
			color = "#192029"
		}

		sq, _ := helpers.GetImageUrl("https://dummyimage.com/360x360/"+strings.ReplaceAll(color, "#", "")+"/000000.png&text=+", 360, 360)

		pfpc.DrawImage(sq, 0, 0)
		pfpc.DrawImage(im, 0, 0)

		dc.DrawImage(pfpc.Image(), 620, 276)
	}

	if err := dc.LoadFontFace("./fonts/coolvetica.ttf", NAME_SIZE); err != nil {
		return nil, err
	}

	for _, dim := range FIELD_POSITIONS {
		dc.DrawStringAnchored(dim.Name, dim.X, dim.Y, 0, 0.5)
	}

	if err := dc.LoadFontFace("./fonts/coolvetica.ttf", VALUE_SIZE); err != nil {
		return nil, err
	}

	for _, dim := range FIELD_POSITIONS {
		dc.DrawStringAnchored(strings.ToUpper(fields[dim.Name]), EOP, dim.Y, 1.0, 0.5)
	}

	if seat, ok := fields["seat"]; ok {
		dc.SetHexColor("#000")
		dc.DrawStringAnchored(strings.ToUpper(seat), 1570.0, 1106.0, 1.0, 0.5)
	}

	qr, _ := qrcode.Encode(fields["qr"], qrcode.Medium, 340)
	if qr != nil {
		reader := bytes.NewReader(qr)
		img, _, _ := image.Decode(reader)
		dc.DrawImage(img, 130, 285)
	}

	return dc.Image(), nil
}

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
