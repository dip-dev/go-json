package opt

import (
	"bytes"
	"encoding/json"
)

// Uint32 ...
type Uint32 struct {
	v      *uint32
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Uint32) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v uint32
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Uint32) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Uint32) Value() *uint32 {
	return opt.v
}

// HasKey ...
func (opt Uint32) HasKey() bool {
	return opt.hasKey
}
