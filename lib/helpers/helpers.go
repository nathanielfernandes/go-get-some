package helpers

import (
	"errors"
	"fmt"
	"image"
	"os"

	"net/http"

	"github.com/nfnt/resize"
)

func GetImage(fp string) image.Image {
	f, err := os.Open(fp)
	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}
	defer f.Close()

	im, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	return im
}

func GetImageUrl(url string, w, h uint) (image.Image, error) {
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
	return GetImageUrl("https://cdn.discordapp.com/emojis/"+e+".png", s, s)
}
