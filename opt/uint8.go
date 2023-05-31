package opt

import (
	"bytes"
	"encoding/json"
)

// Uint8 ...
type Uint8 struct {
	v      *uint8
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Uint8) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v uint8
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Uint8) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Uint8) Value() *uint8 {
	return opt.v
}

// HasKey ...
func (opt Uint8) HasKey() bool {
	return opt.hasKey
}
