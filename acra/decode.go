package acra

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ajg/form"
)

// Decoder decodes ACRA data from a json or x-www-form-urlencoded
type Decoder struct {
	r io.Reader
}

// NewDecoder returns a new Decoder
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r}
}

// Decode reads in and decodes data into dst
func (d Decoder) Decode(dst interface{}) error {
	d1 := json.NewDecoder(d.r)

	err1 := d1.Decode(dst)
	if err1 != nil {
		d2 := form.NewDecoder(d.r)

		err2 := d2.Decode(dst)
		if err2 != nil {
			return fmt.Errorf("%v; %v", err1, err2)
		}
	}

	return nil
}
