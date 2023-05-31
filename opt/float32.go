package opt

import (
	"bytes"
	"encoding/json"
)

// Float32 ...
type Float32 struct {
	v      *float32
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Float32) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v float32
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Float32) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Float32) Value() *float32 {
	return opt.v
}

// HasKey ...
func (opt Float32) HasKey() bool {
	return opt.hasKey
}
