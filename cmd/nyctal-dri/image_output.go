package main

import (
	"fmt"

	"image/jpeg"
	"os"
	"time"

	"nyctal/model"
)

type ImageOutput struct {
	base   string
	frame  int
	period time.Duration
	last   time.Time
}

func NewImageOutput(base string, period time.Duration) model.Output {
	return &ImageOutput{base: base, period: period}
}

func (im *ImageOutput) RenderBuffer(img *model.BGRA) error {
	if time.Since(im.last) > im.period {
		im.last = time.Now()
		im.frame += 1

		outFile, err := os.Create(fmt.Sprintf("%s%03d.jpeg", im.base, im.frame))
		if err != nil {
			return err
		}
		defer outFile.Close()
		err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 100})
		return err
	}
	return nil
}
