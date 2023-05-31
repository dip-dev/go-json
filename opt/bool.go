package opt

import (
	"bytes"
	"encoding/json"
)

// Bool ...
type Bool struct {
	v      *bool
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Bool) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v bool
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Bool) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Bool) Value() *bool {
	return opt.v
}

// HasKey ...
func (opt Bool) HasKey() bool {
	return opt.hasKey
}
