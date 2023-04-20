package utils

import (
	"image/png"
	"math/rand"
	"os"

	"github.com/aofei/cameron"
)

func GenerateImage(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, cameron.Identicon([]byte{byte(rand.Intn(255))}, rand.Intn(512), rand.Intn(512)))
	if err != nil {
		return err
	}

	return nil
}
