package acra

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestEncode(t *testing.T) {
	for _, fn := range files {
		f, err := os.Open(filepath.Join("..", "testdata", fn))
		if err != nil {
			t.Error(err)
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Error(err)
		}

		d := NewDecoder(f)
		r := &Report{}

		err = d.Decode(r)
		if err != nil {
			t.Error(err)
		}

		buf := bytes.NewBuffer(b)
		e := NewEncoder(buf)

		err = e.Encode(r)
		if err != nil {
			t.Error(err)
		}

		f.Close()
	}
}
