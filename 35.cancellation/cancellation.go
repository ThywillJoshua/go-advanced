package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"learn-go/images"
	"log"
	"os"
	"strings"

	"github.com/chai2010/webp"
	"github.com/google/uuid"
)

// load data into pipeline
func makeWork(base64Images ...string) <-chan string {
	out := make(chan string)

	go func() {
		for _, encodedImg := range base64Images {
			out <- encodedImg
		}

		close(out)
	}()

	return out
}

func pipeline[I any, O any](quit <-chan struct{}, input <-chan I, process func(I) O) <-chan O {
	out := make(chan O)
	go func() {
		defer close(out)
		for in := range input {
			select {
			case out <- process(in):
			case <-quit:
				return
			}
		}
	}()
	return out
}

// decode base64 into image format
func base64ToRawImage(base64Img string) image.Image {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Img))

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

// encode as webp
func encodeToWebp(img image.Image) bytes.Buffer {
	var buf bytes.Buffer
	if err := webp.Encode(&buf, img, &webp.Options{Lossless: true}); err != nil {
		log.Fatal(err)
	}

	return buf
}

// save image to disk
func saveToDisk(imgBuf bytes.Buffer) string {
	filename := fmt.Sprintf("%v.webp", uuid.New().String())
	os.WriteFile(filename, imgBuf.Bytes(), 0644)

	return filename
}

func main() {
	base64Images := makeWork(images.Img1, images.Img2, images.Img3)

	quit := make(chan struct{})
	var signal struct{}

	rawImages := pipeline(quit, base64Images, base64ToRawImage)
	webpImages := pipeline(quit, rawImages, encodeToWebp)

	quit <- signal

	filenames := pipeline(quit, webpImages, saveToDisk)

	for name := range filenames {
		fmt.Println(name)
	}
}
