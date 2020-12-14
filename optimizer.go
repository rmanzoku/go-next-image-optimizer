package optimizer

import (
	"context"
	"io/ioutil"
	"net/http"
)

type Optimizer struct {
	SrcURL string
}

func NewOptimizer(srcUrl string) *Optimizer {
	return &Optimizer{SrcURL: srcUrl}
}

func (o *Optimizer) Optimize(url string, w, h, q int, webp bool) ([]byte, error) {
	return o.getSrc(url)
}

func (o *Optimizer) getSrc(url string) ([]byte, error) {
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

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return buf, err
}
