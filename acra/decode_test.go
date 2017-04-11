package acra

import (
	"os"
	"path/filepath"
	"testing"
)

var files = []string{
	"acra-report-android4.2.2.form",
	"acra-report-android4.2.2.json",
}

func TestDecode(t *testing.T) {
	for _, fn := range files {
		f, err := os.Open(filepath.Join("..", "testdata", fn))
		if err != nil {
			t.Error(err)
		}

		d := NewDecoder(f)
		r := &Report{}

		err = d.Decode(r)
		if err != nil {
			t.Error(err)
		}

		f.Close()
	}
}
