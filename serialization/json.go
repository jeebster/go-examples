package serialization

import (
	"encoding/json"
	"io"
)

// ToJSON serializes the interface into a JSON string
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON deserializes the object from a JSON string in the io.Reader to the interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
