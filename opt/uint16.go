package opt

import (
	"bytes"
	"encoding/json"
)

// Uint16 ...
type Uint16 struct {
	v      *uint16
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Uint16) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v uint16
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Uint16) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Uint16) Value() *uint16 {
	return opt.v
}

// HasKey ...
func (opt Uint16) HasKey() bool {
	return opt.hasKey
}
