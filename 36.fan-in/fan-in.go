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
	"sync"

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

func pipeline[I any, O any](input <-chan I, process func(I) O) <-chan O {
	out := make(chan O)
	go func() {
		for in := range input {
			out <- process(in)
		}
		close(out)
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

func fanIn[T any](channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup

	out := make(chan T)

	wg.Add(len(channels))

	for _, ch := range channels {
		go func(in <-chan T) {
			for i := range in {
				out <- i
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	base64Images := makeWork(images.Img1, images.Img2, images.Img3)

	// stage 1
	rawImage1 := pipeline(base64Images, base64ToRawImage)
	rawImage2 := pipeline(base64Images, base64ToRawImage)
	rawImage3 := pipeline(base64Images, base64ToRawImage)

	rawImages := fanIn(rawImage1, rawImage2, rawImage3)

	// stage2:
	webpImages1 := pipeline(rawImages, encodeToWebp)
	webpImages2 := pipeline(rawImages, encodeToWebp)
	webpImages3 := pipeline(rawImages, encodeToWebp)

	webpImages := fanIn(webpImages1, webpImages2, webpImages3)

	// stage3:
	filenames1 := pipeline(webpImages, saveToDisk)
	filenames2 := pipeline(webpImages, saveToDisk)
	filenames3 := pipeline(webpImages, saveToDisk)

	filenames := fanIn(filenames1, filenames2, filenames3)

	for name := range filenames {
		fmt.Println(name)
	}
}
