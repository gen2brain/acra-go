package acra

import (
	"encoding/json"
	"io"
)

// NewEncoder returns a new Encoder
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w}
}

// Encoder encodes ACRA data to json
type Encoder struct {
	w io.Writer
}

// Encode encodes dst as json and writes it out
func (e Encoder) Encode(dst interface{}) error {
	b, err := json.MarshalIndent(dst, "", "    ")
	if err != nil {
		return err
	}

	_, err = e.w.Write(b)
	if err != nil {
		return err
	}

	return nil
}
