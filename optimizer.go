package optimizer

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"net/http"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

type Optimizer struct {
	SrcURL string
}

func NewOptimizer(srcUrl string) *Optimizer {
	return &Optimizer{SrcURL: srcUrl}
}

func (o *Optimizer) Optimize(url string, w, h, q int, webpFlag bool) ([]byte, error) {
	src, err := o.getSrc(url)
	if err != nil {
		return nil, err
	}

	resized := src
	if !(w == 0 && h == 0) {
		resized = imaging.Resize(src, w, h, imaging.NearestNeighbor)
	}

	ret := bytes.Buffer{}
	if webpFlag {
		opt := &webp.Options{
			Lossless: false,
			Quality:  float32(q),
		}
		err = webp.Encode(&ret, resized, opt)
		if err != nil {
			return nil, err
		}
	} else {
		err = png.Encode(&ret, resized)
		if err != nil {
			return nil, err
		}
	}
	return ret.Bytes(), nil
}

func (o *Optimizer) getSrc(url string) (image.Image, error) {
	client := new(http.Client)
	req, err := http.NewRequestWithContext(context.TODO(), "GET", o.SrcURL+url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	return img, err
}
