package optimizer

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	o = &Optimizer{}
)

func TestOptimize(t *testing.T) {
	url := "/images/heroes/2000/5001.png"
	w := 64
	h := 64
	q := 100
	ret, err := o.Optimize(url, w, h, q, true)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile("test.png", ret, 0777)
	if err != nil {
		t.Fatal(err)
	}
}
func TestMain(m *testing.M) {
	o = NewOptimizer("https://www.mycryptoheroes.net")
	code := m.Run()
	os.Exit(code)
}
