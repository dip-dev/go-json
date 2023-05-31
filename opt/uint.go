package opt

import (
	"bytes"
	"encoding/json"
)

// Uint ...
type Uint struct {
	v      *uint
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Uint) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v uint
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Uint) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Uint) Value() *uint {
	return opt.v
}

// HasKey ...
func (opt Uint) HasKey() bool {
	return opt.hasKey
}
