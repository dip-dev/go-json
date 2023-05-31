package opt

import (
	"bytes"
	"encoding/json"
)

// Uint64 ...
type Uint64 struct {
	v      *uint64
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Uint64) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v uint64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Uint64) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Uint64) Value() *uint64 {
	return opt.v
}

// HasKey ...
func (opt Uint64) HasKey() bool {
	return opt.hasKey
}
