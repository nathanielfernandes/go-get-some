package helpers

import (
	"errors"
	"image"

	"net/http"

	"github.com/nfnt/resize"
)

func GetImage(url string, w, h uint) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Non 200 status")
	}
	im, _, err := image.Decode(resp.Body)

	if err == nil {
		im = resize.Resize(w, h, im, resize.Bilinear)
	}
	return im, err
}

func GetEmoji(e string, s uint) (image.Image, error) {
	return GetImage("https://cdn.discordapp.com/emojis/"+e+".png", s, s)
}
